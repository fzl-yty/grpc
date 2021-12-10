package regulator

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"git.xiaojukeji.com/nuwa/golibs/coordinator/consts"
	"git.xiaojukeji.com/nuwa/golibs/coordinator/hash"
	"git.xiaojukeji.com/nuwa/golibs/coordinator/job"
	"git.xiaojukeji.com/nuwa/golibs/coordinator/proto"
	"git.xiaojukeji.com/nuwa/golibs/coordinator/utils"

	"strings"
	"time"

	"git.xiaojukeji.com/nuwa/golibs/coordinator/log"
	zK "git.xiaojukeji.com/nuwa/golibs/coordinator/zk"
)

type Event struct {
	JobName string
	Mode    string
	Key     []string
	Timeout time.Duration
	Retry   int
}

type Coordinator struct {
	leaderIP string

	electionWatcher *zK.ProcessNode
	node2Job        map[string]string
	nodeStopCh      <-chan struct{}

	GRPCClient
	CordGRPCServer

	EventCh  chan Event // 业务触发event chan
	CustomCh chan Event // 自定义模式chan
	ErrCh    chan error // 任务执行错误chan

	// 本机IP、hostName
	ip       string
	hostName string

	// 一致性hash
	hashRing *hash.HashRing

	*option
}

var defaultCoordinator *Coordinator

func NewCoordinator(opts ...Option) (*Coordinator, error) {

	opt := &option{}

	for _, o := range opts {
		o(opt)
	}

	if len(opt.zkIPs) == 0 || opt.grpcPort == "" {
		return nil, errors.New("config params err")
	}

	if opt.commonLog != nil {
		log.CLog = opt.commonLog
	} else {
		log.CLog = &log.CLogHandle{}
	}

	if defaultCoordinator != nil {
		return defaultCoordinator, nil
	}
	IP := utils.GetIP()
	if IP == "" {
		return nil, errors.New("get local IP err")
	}

	hostName, err := utils.GetHostName()
	if err != nil {
		return nil, err
	}

	// zK相关设置
	nodeErrCh := make(chan error)
	nodeElectedCh := make(chan bool)
	nodeStopCh := make(<-chan struct{})

	nodeProcess, err := zK.NewProcessNode(IP, opt.zkIPs, nodeErrCh, nodeElectedCh, log.CLog)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
			"register zK new ProcessNode err || errMsg=%s", err.Error())
		return nil, err
	}

	eventCh := make(chan Event, 10)
	errCh := make(chan error, 10)
	//构造RPC server
	defaultCoordinator = &Coordinator{
		nodeStopCh:      nodeStopCh,
		electionWatcher: nodeProcess,
		EventCh:         eventCh,
		ErrCh:           errCh,
		ip:              IP,
		hostName:        hostName,
		option:          opt,
	}
	leaderCh := make(chan *proto.JobKeyReq, 10)
	workerCh := make(chan *proto.ChildJobReq, 10)
	RPCServer := NewGRPCServer(defaultCoordinator, leaderCh, workerCh)
	defaultCoordinator.CordGRPCServer = RPCServer
	defaultCoordinator.GRPCClient = NewGRPCClient(nil)

	return defaultCoordinator, nil
}

// Coordinator需要做的事情
func (c *Coordinator) Run() error {
	// 启动zK监听
	_, err := c.electionWatcher.ZkService.WatchTree(consts.WorkerNode, c.nodeStopCh, nil)
	if err != nil {
		return err
	}
	err = c.electionWatcher.Run(consts.WorkerNode)
	if err != nil {
		return err
	}

	// 注册当前节点信息
	curNode := &zK.Node{NodeType: "worker", NodeName: c.hostName, NodeIP: c.ip}
	b, err := json.Marshal(curNode)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
			"json marshal node msg err || errMsg=%s", err.Error())
	}
	key := strings.TrimLeft(c.electionWatcher.ProcessNodePath, "/")
	err = c.electionWatcher.ZkService.Put(key, b, nil)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
			"put zK node msg err || errMsg=%s", err.Error())
	}

	// 启动grpc服务
	go c.serve()

	// 监听业务事件
	go func(cord *Coordinator) {
		for {
			select {
			// node选举为leader
			case e := <-cord.electionWatcher.ElectedCh:
				// 注册leader
				if e {
					err = cord.electionWatcher.ZkService.Put(consts.LeaderNode, b, nil)
					if err != nil {
						fmt.Println("zKService.Put err", err.Error())
					}
					log.CLog.Infof(context.Background(), consts.CoordinatorPrefix,
						"%s is elected as leader", cord.ip)
					// 获取10s内未执行完成的任务
					err := cord.reDistributionByTime(cord.checkByTime)
					if err != nil {
						log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
							"校验zK为任务失败 || errMsg=%s", err.Error())
					}
				}
			// leader获取任务，开始执行
			case leaderJobEvent := <-cord.LeaderCh:
				log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "leader get jobKey=%s||jobMode=%s||jobNode=%s",
					leaderJobEvent.Name, leaderJobEvent.Mode, leaderJobEvent.JobNode)

				// 构造任务, 更新任务时间信息
				now := time.Now()
				var newJob = job.Job{
					WorkerTriggerTime:     leaderJobEvent.WorkerTriggerTime.AsTime(),
					LeaderTriggerTime:     now,
					Name:                  leaderJobEvent.Name,
					Mode:                  leaderJobEvent.Mode,
					Keys:                  leaderJobEvent.Keys,
					DistributedKeys:       make(map[string][]string, 0),
					DistributedFailedKeys: make(map[string][]string, 0),
					AllocationResult:      false,
					Node:                  leaderJobEvent.JobNode,
				}

				// 创建任务节点,主节点收到job信息,上报zK信息
				b, err := json.Marshal(newJob)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
						"leader recv job but json Marshal job err || errMsg=%s", err.Error())
				}
				err = c.electionWatcher.ZkService.Put(newJob.Node, b, nil)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
						"leader recv job but update job to zK err|| errMsg=%s", err.Error())
				}
				err = cord.doJob(&newJob)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix,
						"leader Do job err || errMsg=%s", err.Error())
				}

			// worker接收到业务任务，上次zK，并通知leader
			case bizEvent := <-cord.EventCh:
				// 单机测试时候,需要sleep,等待注册好leader节点
				time.Sleep(time.Second)
				// 构造上报zK的job信息
				now := time.Now()
				JobName := bizEvent.JobName + "_" + strconv.Itoa(int(now.Unix()))
				newJob := &job.Job{
					WorkerTriggerTime: now,
					UpdateTime:        now,
					Name:              JobName,
					Mode:              bizEvent.Mode,
					Keys:              bizEvent.Key,
					AllocationResult:  false,
				}
				newJob.Node = consts.JobState + "/" + newJob.Name

				// 创建任务节点
				c.electionWatcher.ZkService.CreateNode(newJob.Node, false, false)
				b, err := json.Marshal(newJob)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "worker json marshal job err||errMsg=%s",
						err.Error())
				}

				// 首次登记任务信息由worker发起, 第一份数据流错误zKErr
				zKErr := c.electionWatcher.ZkService.Put(newJob.Node, b, nil)
				if zKErr != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "msg=worker put jobMsg to zK err||errMsg=%s",
						zKErr.Error())
				}

				// leaderErr，第二份数据流错误
				leader, leaderErr := c.electionWatcher.GetLeader()
				leaderAddr := leader.NodeIP + ":" + c.grpcPort
				log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "msg=worker get job||jobKey=%s||jobMode=%s||jobName=%s||leaderAddr=%s",
					bizEvent.Key, bizEvent.Mode, bizEvent.JobName, leaderAddr)
				if err != nil {
					//获取leader信息失败, 不能直接返回, 将任务信息上报至zK, 等待新主节点获取任务执行, 此处只打印错误日志
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "msg=worker get leaderAddr err||errMsg=%s",
						err.Error())
				}
				leaderErr = c.notifyLeaderTaskStart(context.Background(), leaderAddr, newJob)
				if leaderErr != nil {
					// 写到errChan中, 通知业务任务执行失败
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "msg=worker NotifyLeader jobMsg err||errMsg=%s",
						leaderErr.Error())
				}
				// 如果两份数据传递都失败,则说明该节点被孤立了，返回业务错误
				if leaderErr != nil && zKErr != nil {
					errMsg := fmt.Errorf("msg=任务执行失败||jobName=%s||send zK errMsg=%s||otifiedLeaderErr=%s", newJob.Name, zKErr.Error(), leaderErr.Error())
					cord.ErrCh <- errMsg
				}
			}
		}
	}(c)
	// 判断上一次任务执行状态
	return nil
}

func (c *Coordinator) broadcastFunc(job *job.Job) error {
	// 获取所有节点
	var (
		nodes []*zK.Node
		err   error
	)
	nodes, err = c.electionWatcher.GetNodes(consts.WorkerNode)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "Broadcast get nodes err||errMsg=%s",
			err.Error())
		return err
	}
	log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "do broadcastFunc||job name=%s||job mode=%s||job keys=%v",
		job.Name, job.Mode, job.Keys)

	// 广播任务到子节点
	for _, v := range nodes {
		// 此处记录子任务状态未分发，存入store
		addr := v.NodeIP + ":" + c.grpcPort
		err := c.distributeSubTasks(context.Background(), addr, job.Mode, job.Name, job.Keys)
		if err != nil {
			job.DistributedFailedKeys[addr] = job.Keys
			b, err := json.Marshal(job)
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "Broadcast to node err||errMsg=%s",
				err)
			err = c.electionWatcher.ZkService.Put(job.Node, b, nil)
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "update broadcast to zK node err||errMsg=%s",
				err)
			continue
		}

		job.DistributedKeys[addr] = job.Keys
		log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "Broadcast to node ||worker addr=%s||worker keys=%v",
			addr, job.Keys)
		b, err := json.Marshal(job)
		log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "Broadcast to node, ip=%s||keys=%v", addr, job.Keys)
		err = c.electionWatcher.ZkService.Put(job.Node, b, nil)
	}

	go c.clearJobRecord(24 * time.Hour)
	return nil
}

func (c *Coordinator) hashFunc(job *job.Job) error {
	err := c.updateHashRing()
	if err != nil {
		// 执行失败不返回，使用原有的hashRing进行分片
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "do hashJob but updateHashRing err ||errMsg=%s",
			err.Error())
	}
	// mapKeys记录成功分配的ip
	var mapKeys = make(map[string][]string, 0)
	for _, v := range job.Keys {
		childKey := fmt.Sprintf("%v", v)
		nodeIP := c.hashRing.GetNode(childKey) + ":" + c.grpcPort
		mapKeys[nodeIP] = append(mapKeys[nodeIP], childKey)
	}

	for addr, v := range mapKeys {
		err := c.distributeSubTasks(context.Background(), addr, consts.HashMode, job.Name, v)
		if err != nil {
			// 记录发布失败的keys
			job.DistributedFailedKeys[addr] = v
			b, err := json.Marshal(job)
			if err != nil {
				log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "json marshal err but updateHashRing err ||errMsg=%s",
					err.Error())
			}
			err = c.electionWatcher.ZkService.Put(job.Node, b, nil)
			if err != nil {
				log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "update zK job status err ||errMsg=%s",
					err.Error())
			}
			// 移除失败的节点
			c.hashRing.RemoveNode(addr)
			delete(mapKeys, addr)
			continue
		}
		job.DistributedKeys[addr] = v
		b, err := json.Marshal(job)
		if err != nil {
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "json marshal err but updateHashRing err ||errMsg=%s",
				err.Error())
		}
		err = c.electionWatcher.ZkService.Put(job.Node, b, nil)
		if err != nil {
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "update zK job status err ||errMsg=%s",
				err.Error())
		}
	}
	// remainIPs包含分配成功的IP或者未被分配的IP
	var remainIPs = make([]string, 0)
	for i := range mapKeys {
		remainIPs = append(remainIPs, i)
	}

	// 将未分配的keys, 在成功的IP中随机选择一个执行
	if len(job.DistributedFailedKeys) > 0 {
		length := len(mapKeys)
		// 随机从剩余的ip中选择一个
		for oldIP, keys := range job.DistributedFailedKeys {
			for i := 0; i < 1; i++ {
				ipIndex := rand.Intn(length)
				remainKs := keys
				NewIP := remainIPs[ipIndex]
				err := c.distributeSubTasks(context.Background(), NewIP, consts.HashMode, job.Name, remainKs)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "重新分配子任务失败||jobName=%s||newAddr=%s||subKeys=%v||errMsg=%s",
						job.Name, NewIP, remainKs, err.Error())
					continue
				}
				log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "重新分配子任务成功||jobName=%s||newAddr=%s||subKeys=%v",
					job.Name, NewIP, remainKs)
				// 更新job信息,重新分配成功
				delete(job.DistributedFailedKeys, oldIP)
				job.DistributedKeys[NewIP] = append(job.DistributedKeys[NewIP])

				b, err := json.Marshal(job)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "json marshal err ||errMsg=%s",
						err.Error())
				}
				err = c.electionWatcher.ZkService.Put(job.Node, b, nil)
				if err != nil {
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "重新分配子任务状态更新失败 ||errMsg=%s",
						err.Error())
				}
			}
		}
	}

	go c.clearJobRecord(24 * time.Hour)
	return nil
}

// 只有leader处理
func (c *Coordinator) customFunc(job *job.Job) error {
	var event = Event{
		Mode:    job.Mode,
		Key:     job.Keys,
		Timeout: 0,
		Retry:   0,
	}
	select {
	case c.CustomCh <- event:
	default:
	}
	go c.clearJobRecord(24 * time.Hour)
	return nil
}

func (c *Coordinator) clearJobRecord(keepTime time.Duration) {
	now := time.Now().Unix()
	children := c.electionWatcher.ZkService.GetChildren(consts.JobState, false)
	for _, v := range children {
		CreateTimeStr := strings.Split(v, "_")[1]
		CreateTime, _ := strconv.Atoi(CreateTimeStr)
		if now-int64(CreateTime) > int64(keepTime)/1e9 {
			_ = c.electionWatcher.ZkService.Delete(v)
		}
	}
	return
}

func (c *Coordinator) doJob(job *job.Job) error {
	switch job.Mode {
	case consts.BroadcastMode:
		return c.broadcastFunc(job)
	case consts.HashMode:
		return c.hashFunc(job)
	// 后续扩展
	//case consts.CustomMode:
	//	return c.customFunc(job)
	default:
		return nil
	}
}

// 主节点更换后, 检查在时间范围t之内的任务执行情况, 执行失败的任务会继续执行
func (c *Coordinator) reDistributionByTime(t time.Duration) error {
	//获取当前时间节点
	now := time.Now().Unix()
	duration := t / 1e9
	children := c.electionWatcher.ZkService.GetChildren(consts.JobState, false)

	// 节点变更, 重新生成hashRing
	for _, childPath := range children {
		CreateTimeStr := strings.Split(childPath, "_")[1]
		CreateTime, _ := strconv.Atoi(CreateTimeStr)
		if now - int64(CreateTime) < int64(duration) {
			log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "主节点获取未执行完成的任务||jobPath=%s",
				childPath)
			jobKV, err := c.electionWatcher.ZkService.Get(childPath, nil)
			if err != nil {
				log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "主节点获取任务信息失败||jobPath=%s||errMsg=%s",
					childPath, err.Error())
				continue
			}
			var oldJob = new(job.Job)
			err = json.Unmarshal(jobKV.Value, oldJob)
			if err != nil {
				log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "ReDistributionByTime json marshal err ||errMsg=%s",
					err.Error())
				continue
			}
			// 根据模式不同，选择不同的恢复模式
			switch oldJob.Mode {
			case consts.BroadcastMode:
				err := c.recoverBroadcastJob(oldJob)
				if err != nil{
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "recoverBroadcastJob err ||errMsg=%s",
						err.Error())
				}
			case consts.HashMode:
				err := c.recoverHashJob(oldJob)
				if err != nil{
					log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "recoverHashJob err ||errMsg=%s",
						err.Error())
				}
			default:
				log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "ReDistributionByTime unKnown mode ||mode=%s", oldJob.Mode)
			}
		}
	}
	return nil
}

func (c *Coordinator) updateHashRing() error {
	nodes, err := c.electionWatcher.GetNodes(consts.WorkerNode)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "updateHashRing err || errMsg=%s", err.Error())
		return err
	}
	if len(nodes) == 0 {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "updateHashRing err || errMsg=no nodes available", "")
		return errors.New("no nodes available")
	}
	// 构造hash环
	nodeWeight := make(map[string]int)
	for _, v := range nodes {
		nodeWeight[v.NodeIP] = consts.HashWeight
	}
	c.hashRing = hash.NewHashRing(-1)
	c.hashRing.AddNodes(nodeWeight)
	return nil
}

// 恢复hash任务
func (c *Coordinator) recoverHashJob(oldJob *job.Job) error {
	err := c.updateHashRing()
	if err != nil {
		return err
	}
	// 以IP上次分配失败的节点IP做新的key，重新分配
	for OldIP, failedKeys := range oldJob.DistributedFailedKeys {
		keys := failedKeys
		// 重复一次
		var NewIP string
		var err error
		for i := 0; i < 2; i++ {
			NewIP = c.hashRing.GetNode(OldIP) + ":" + c.grpcPort
			if NewIP == "" {
				return errors.New("hashRing get node err||errMsg=no nodes available")
			}

			err = c.distributeSubTasks(context.Background(), NewIP, oldJob.Node, oldJob.Name, keys)
			if err != nil {
				// 新节点与当前leader连接有问题，删除该节点，重新分配一次
				c.hashRing.RemoveNode(NewIP)
				log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "主节点变更重新分配子任务失败,jobName=%s||newIP=%s||keys=%v", oldJob.Name, NewIP, keys)
				continue
			}
			// 分配成功直接退出本次循环
			break
		}
		// 重新分配两次依然失败（概率很低），则该keys不再分配，分配其它失败的keys
		if err != nil {
			continue
		}

		// 更新任务信息
		oldJob.DistributedKeys[NewIP] = append(oldJob.DistributedKeys[NewIP], oldJob.DistributedFailedKeys[OldIP]...)
		delete(oldJob.DistributedFailedKeys, OldIP)
		data, err := json.Marshal(&oldJob)
		if err != nil {
			// 更新失败
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "主节点变更重新分配子任务失败,jobName=%s||newIP=%s||keys=%v", oldJob.Name, NewIP, keys)
			continue
		}
		err = c.electionWatcher.ZkService.Put(oldJob.Node, data, nil)
		if err != nil {
			// 更新失败
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "主节点变更重新分配子任务失败,jobName=%s||newIP=%s||keys=%v", oldJob.Name, NewIP, keys)
		}
	}
	return nil
}

// 恢复广播任务
func(c *Coordinator) recoverBroadcastJob(oldJob *job.Job) error{
	nodes, err := c.electionWatcher.GetNodes(consts.WorkerNode)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "Broadcast get nodes err||errMsg=%s",
			err.Error())
		return err
	}

	for _, v := range nodes {
		nodeAddr := v.NodeIP + ":" + c.grpcPort
		_, ok := oldJob.DistributedFailedKeys[nodeAddr]
		if ok {
			continue
		}
		err := c.distributeSubTasks(context.Background(), nodeAddr, oldJob.Mode, oldJob.Name, oldJob.Keys)
		// 广播失败, 此处不重试
		if err != nil {
			oldJob.DistributedFailedKeys[nodeAddr] = oldJob.Keys
			b, err := json.Marshal(oldJob)
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "Broadcast to node err||errMsg=%s",
				err)
			err = c.electionWatcher.ZkService.Put(oldJob.Node, b, nil)
			log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "update broadcast to zK node err||errMsg=%s",
				err)
			continue
		}

		oldJob.DistributedKeys[nodeAddr] = oldJob.Keys
		log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "Broadcast to node ||worker addr=%s||worker keys=%v",
			nodeAddr, oldJob.Keys)
		b, err := json.Marshal(oldJob)
		log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "Broadcast to node, ip=%s||keys=%v", nodeAddr, oldJob.Keys)
		err = c.electionWatcher.ZkService.Put(oldJob.Node, b, nil)
	}

	return nil
}
