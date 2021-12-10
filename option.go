package regulator

import (
	"git.xiaojukeji.com/nuwa/golibs/coordinator/log"
	"time"
)

type option struct {
	grpcPort  string
	commonLog log.CommonLog
	zkIPs     []string

	// checkByTime检查该时间范围内未执行的任务重新执行
	checkByTime time.Duration
}

type Option func(o *option)

func SetGrpcPort(port string) Option {
	return func(o *option) {
		o.grpcPort = port
	}
}

func SetLog(clog log.CommonLog) Option {
	return func(o *option) {
		o.commonLog = clog
	}
}

func SetBackends(IPs []string) Option {
	return func(o *option) {
		o.zkIPs = IPs
	}
}

func SetCheckByTime(t time.Duration) Option {
	return func(o *option) {
		o.checkByTime = t
	}
}


