package main

import (
	"log"
	"net"

	kServer "github.com/cloudwego/kitex/server"

	"kitex-example/internal/interfaces/rpc"
	pb "kitex-example/internal/pb/greeter"
)

func main() {
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
