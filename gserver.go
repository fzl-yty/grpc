package regulator

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/fzl-yty/zkElection/consts"
	"github.com/fzl-yty/zkElection/log"
	"github.com/fzl-yty/zkElection/proto"
	"github.com/fzl-yty/zkElection/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CordGRPCServer struct {
	// worker通知leader启动任务的channel
	LeaderCh chan *proto.JobKeyReq
	// worker接收leader下发任务channel
	WorkerCh chan *proto.ChildJobReq
	cord     *Coordinator
}

func NewGRPCServer(c *Coordinator, leaderCh chan *proto.JobKeyReq, workerCh chan *proto.ChildJobReq) CordGRPCServer {
	return CordGRPCServer{
		cord:     c,
		LeaderCh: leaderCh,
		WorkerCh: workerCh,
	}
}

func (c *CordGRPCServer)serve() {
	localIP := utils.GetIP()
	addr := localIP + ":" + c.cord.grpcPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "", err.Error())
		fmt.Println("grpc serve err:", err.Error())
		return
	}
	s := grpc.NewServer()
	proto.RegisterJobServiceServer(s, c)
	reflection.Register(s)
	go s.Serve(lis)
	return
}

func (c *CordGRPCServer) TaskStart(ctx context.Context, in *proto.JobKeyReq) (*proto.JobKeyResp, error) {
	select {
	case c.LeaderCh <- in:
		return &proto.JobKeyResp{
			Data:       "",
			Errno:      "",
			ErrMessage: fmt.Sprintf("%s is stored in the leader execution queue", in.Name),
		}, nil
	default:
		return &proto.JobKeyResp{
			Data:       "",
			Errno:      "10000",
			ErrMessage: fmt.Sprintf("send %s to leader failed", in.Name),
		}, errors.New("")
	}
}

func (c *CordGRPCServer) SubTaskStart(ctx context.Context, in *proto.ChildJobReq) (*proto.ChildJobResp, error) {
	select {
	case c.cord.WorkerCh <- in:
		log.CLog.Infof(context.Background(), consts.CoordinatorPrefix, "msg=worker get Sub keys||job name=%s||job mode=%s||keys=%s",
			in.JobName, in.Mode, in.JobKeys)
		return &proto.ChildJobResp{
			Data:       "",
			Errno:      "",
			ErrMessage: fmt.Sprintf("%s is stored in the leader execution queue", in.JobName),
		}, nil
	default:
		log.CLog.Errorf(context.Background(), consts.CoordinatorPrefix, "msg=worker get Sub keys err||errMsg=put sub job to WorkerCh err")
		return &proto.ChildJobResp{
			Data:       "",
			Errno:      "10000",
			ErrMessage: fmt.Sprintf("send %s to worker failed", in.JobName),
		}, errors.New("")
	}
}
