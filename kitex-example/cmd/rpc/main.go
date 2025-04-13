package main

import (
	"log"
	"net"
	"time"

	kServer "github.com/cloudwego/kitex/server"
	prometheus "github.com/kitex-contrib/monitor-prometheus"

	"kitex-example/internal/infras/logger"
	"kitex-example/internal/interfaces/rpc"
	"kitex-example/internal/interfaces/rpc/middleware"
	pb "kitex-example/internal/pb/greeter"
)

func main() {
	logger.Default(
		logger.WithStdout(true),
		logger.WithJsonFormat(true),
		logger.WriteToFile(false),
	)

	// 服务运行地址
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8890")
	// 创建服务实例，更多配置参考kitex官方文档
	svr := pb.NewServer(
		new(rpc.GreeterImpl),
		kServer.WithServiceAddr(addr),
		// prometheus接入
		kServer.WithTracer(prometheus.NewServerTracer(":9093", "/metrics")),
		// 日志中间件
		kServer.WithMiddleware(middleware.AccessLog),
		// 参数验证中间件
		kServer.WithMiddleware(middleware.Validator),
		// Server 在收到退出信号时的等待时间
		// 如果超过该等待时间，Server 将会强制结束所有在处理的请求
		kServer.WithExitWaitTime(5*time.Second), // graceful shutdown timeout
		// 在连接上读写数据所能忍受的最大等待时间，主要为防止异常连接卡住协程
		kServer.WithReadWriteTimeout(10*time.Second), // read and write timeout
	)
	err := svr.Run()
	if err != nil {
		log.Printf("run server err: %v", err)
	}
}
