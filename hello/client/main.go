package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	etcd "github.com/kitex-contrib/registry-etcd"

	"hello/kitex_gen/pb"
	"hello/kitex_gen/pb/greeter"
)

// 是否开启服务发现，这里使用的是etcd作为服务发现和注册
var enableDiscovery = false

func main() {
	var (
		c   greeter.Client
		err error
	)

	if enableDiscovery { // 如果开启了服务发现
		var r discovery.Resolver
		r, err = etcd.NewEtcdResolver(
			[]string{"127.0.0.1:12379"},
			etcd.WithEtcdServicePrefix("kitex/registry-etcd"),
		)
		if err != nil {
			log.Fatal("new etcd resolver err:", err)
		}
		c, err = greeter.NewClient("services.greeter", client.WithResolver(r))
	} else {
		c, err = greeter.NewClient("services.greeter", client.WithHostPorts("0.0.0.0:8890"))
	}
	if err != nil {
		log.Fatal("new client err:", err)
	}

	// 调用rpc方法
	for {
		req := &pb.HelloRequest{Msg: "my request"}
		resp, e := c.Hello(context.Background(), req)
		if e != nil {
			log.Fatal("rpc error:", e)
		}

		log.Println(resp.Msg)
		time.Sleep(time.Second)
	}
}
