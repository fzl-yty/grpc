package regulator

import (
	"context"
	"sync"

	"github.com/fzl-yty/zkElection/consts"
	"github.com/fzl-yty/zkElection/job"
	"github.com/fzl-yty/zkElection/log"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/fzl-yty/zkElection/proto"
	"google.golang.org/grpc"
)

type GRPCClient interface {
	connect(string) (*grpc.ClientConn, error)
	distributeSubTasks(ctx context.Context, addr, mode, JobName string, childKey []string) error
	notifyLeaderTaskStart(ctx context.Context, addr string, job *job.Job) error
}

type CordGRPCClient struct {
	dialOpt grpc.DialOption
	wg      sync.WaitGroup
}

func NewGRPCClient(dialOpt grpc.DialOption) GRPCClient {
	if dialOpt == nil {
		dialOpt = grpc.WithInsecure()
	}
	return &CordGRPCClient{dialOpt: dialOpt}
}

func (c *CordGRPCClient) connect(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, c.dialOpt)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (c *CordGRPCClient) notifyLeaderTaskStart(ctx context.Context, addr string, job *job.Job) error {
	var conn *grpc.ClientConn
	conn, err := c.connect(addr)
	defer conn.Close()
	if err != nil {
		return err
	}
	d := proto.NewJobServiceClient(conn)

	var notifiedJob = proto.JobKeyReq{
		WorkerTriggerTime: timestamppb.New(job.WorkerTriggerTime),
		UpdateTime:        timestamppb.New(job.UpdateTime),
		Name:              job.Name,
		Mode:              job.Mode,
		Keys:              job.Keys,
		JobNode:           job.Node,
		AllocationResult:  false,
	}

	_, err = d.TaskStart(ctx, &notifiedJob)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "NotifyLeaderTaskStart err||"+
			"errMsg=%s||job node=%s||job Mode=%s||job name=%s||job keys=%v", err.Error(), job.Node,
			job.Mode, job.Name, job.Keys)
		return err
	}
	log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "NotifyLeaderTaskStart||job node=%s||job Mode=%s||"+
		"job name=%s||job keys=%v", job.Node, job.Mode, job.Name, job.Keys)
	return nil
}

func (c CordGRPCClient) distributeSubTasks(ctx context.Context, addr, mode, jobName string, childKeys []string) error {
	var conn *grpc.ClientConn
	conn, err := c.connect(addr)
	defer conn.Close()
	if err != nil {
		return err
	}
	d := proto.NewJobServiceClient(conn)
	_, err = d.SubTaskStart(ctx, &proto.ChildJobReq{Mode: mode, JobName: jobName, JobKeys: childKeys})
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "distributeSubTasks err||"+
			"errMsg=%s||worker addr=%s||job Mode=%s||job name=%s||sub keys=%v", err.Error(), addr,
			mode, jobName, childKeys)
		return err
	}
	return nil
}
