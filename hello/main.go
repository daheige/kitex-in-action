package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"

	kServer "github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/monitor-prometheus"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"

	pb "hello/kitex_gen/pb/greeter"
)

func main() {
	pid := os.Getpid()
	fmt.Printf("current service pid: %d\n", pid)

	// 日志输出采用zap框架实现日志json格式输出
	klog.SetLogger(kitexzap.NewLogger())
	//klog.SetLevel(klog.LevelDebug)
	klog.SetLevel(klog.LevelInfo)

	// 可以根据实际情况将日志输出到文件中
	//f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	log.Fatal("open output file err:", err)
	//}
	//defer f.Close()
	//klog.SetOutput(f) // 将日式重定向到文件

	// 使用时请传入真实 etcd 的服务地址，本例中为 127.0.0.1:12379
	r, err := etcd.NewEtcdRegistry(
		[]string{"127.0.0.1:12379"}, etcd.WithEtcdServicePrefix("kitex/registry-etcd"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 服务运行地址
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8890")
	svr := pb.NewServer(new(GreeterImpl),
		kServer.WithServiceAddr(addr),
		// 指定 Registry 与服务基本信息
		kServer.WithRegistry(r),
		kServer.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "services.greeter",
			},
		),
		// prometheus接入
		kServer.WithTracer(prometheus.NewServerTracer(":9093", "/metrics")),
	)

	// 启动服务
	err = svr.Run()
	if err != nil {
		log.Println(err)
	}
}
