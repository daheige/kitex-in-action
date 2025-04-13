package main

import (
	"context"
	"log"

	"github.com/cloudwego/kitex/client"

	"kitex-example/internal/pb"
	"kitex-example/internal/pb/greeter"
)

func main() {
	var (
		c   greeter.Client
		err error
	)

	// 如果需要接入服务发现和注册参考 kitex-in-action/hello
	c, err = greeter.NewClient("services.greeter", client.WithHostPorts("0.0.0.0:8890"))
	if err != nil {
		log.Fatal("new client err:", err)
	}

	// 调用rpc方法
	req := &pb.HelloRequest{Msg: ""}
	resp, e := c.Hello(context.Background(), req)
	if e != nil {
		log.Fatal("rpc error:", e)
	}

	log.Println("resp msg:", resp.Msg)
}
