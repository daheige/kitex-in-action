package main

import (
	"context"
	"log"
	"net"

	kServer "github.com/cloudwego/kitex/server"
	"go.uber.org/zap"

	"kitex-example/internal/infras/logger"
	"kitex-example/internal/interfaces/rpc"
	pb "kitex-example/internal/pb/greeter"
)

func main() {
	logger.Default(logger.WithStdout(true), logger.WithJsonFormat(true), logger.WriteToFile(false))
	ctx := context.Background()
	logger.Info(ctx, "hello world", zap.String("foo", "bar"))

	// 服务运行地址
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8890")
	// 创建服务实例，更多配置参考kitex官方文档
	svr := pb.NewServer(
		new(rpc.GreeterImpl),
		kServer.WithServiceAddr(addr),
	)
	err := svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
