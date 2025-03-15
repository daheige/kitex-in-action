# kitex实战
参考文档：https://www.cloudwego.io/zh/docs/kitex/

# 代码生成之前的准备工作
参考[kitex-in-action/readme.md](../readme.md)

# 根据proto文件生成脚手架layout
1. 定义好proto文件，参考 `idl/helloworld.proto` 文件
2. 通过kitex工具生成脚手架代码
```shell
# 创建项目目录
mkdir -p hello/idl
# 仅仅是生成脚手架代码，这里的-module参数是项目的目录名字
kitex -service greeter -module hello -I idl idl/helloworld.proto
```

3. 获取相关的go代码包
```shell
go mod tidy
```
此时代码生成的结构如下所示：
```
├── build.sh
├── client
│   └── main.go
├── go.mod
├── go.sum
├── handler.go
├── idl
├── kitex_gen # kitex脚手架生成的代码目录
│   └── pb
├── kitex_info.yaml
├── main.go
├── output
│   ├── bin
│   ├── bootstrap.sh
│   └── log
├── readme.md
└── script
    ├── bootstrap.sh
```

# 运行项目
1. 编译构建应用程序
```shell
sh build.sh
```
此时就会在当前项目下生成一个output目录，包含bin/greeter二进制文件。这个output目录中的bootstrap.sh文件，可以快速在本地启动项目。

2. 本地运行etcd服务(这里使用的服务发现和注册是etcd，当然你可以根据实际情况选择不同的组件)
```shell
sh start-etcd.sh
```

3. 启动服务
```shell
sh output/bootstrap.sh
```
服务启动后效果如下：
```
2025/03/01 21:50:15.363992 server.go:79: [Info] KITEX: server listen at addr=[::]:8890
2025/03/01 21:50:16.369342 etcd_registry.go:299: [Info] start keepalive lease 694d9551e212d53e for etcd registry
```

# proto文件发生变化后的代码生成
```shell
# 当proto文件发生更改，执行该命令，并实现对应的service方法即可
kitex -module hello -I idl idl/helloworld.proto
```

# 客户端运行验证
执行 go run client/main.go 命令，结果如下：
```shell
2025/03/01 22:16:47 hello,my request
2025/03/01 22:16:48 hello,my request
2025/03/01 22:16:49 hello,my request
2025/03/01 22:16:50 hello,my request
```

# 通过etcd查看注册的服务
```shell
docker exec -it etcd_test /bin/bash
etcdctl get kitex/registry-etcd --prefix
```
输出结果如下：
```
kitex/registry-etcd/services.greeter/192.168.10.101:8890
{"network":"tcp","address":"192.168.10.101:8890","weight":10,"tags":null}
```

# 通过命令行工具请求grpc服务
1. 安装kitexcall工具
```shell
go install github.com/kitex-contrib/kitexcall@latest
```
2. 通过kitexcall请求数据
```shell
# -idl-path 用于指定proto文件 -m用于指定serviceName/method -d请求数据可以使用json格式 -e用于指定服务地址
kitexcall -idl-path idl/helloworld.proto -m Greeter/Hello -d '{"msg": "kitex"}' -e 127.0.0.1:8890
```

# 日志接入
https://www.cloudwego.io/zh/docs/kitex/tutorials/observability/logging/
```go
// 日志输出采用zap框架实现日志json格式输出
klog.SetLogger(kitexzap.NewLogger())
klog.SetLevel(klog.LevelDebug)

// 可以根据实际情况将日志输出到文件中
f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
log.Fatal("open output file err:", err)
}
defer f.Close()
klog.SetOutput(f) // 将日式重定向到文件
```

# 服务发现和注册参考
https://www.cloudwego.io/zh/docs/kitex/tutorials/third-party/service_discovery/etcd/

# 服务可观测性
https://www.cloudwego.io/zh/docs/kitex/tutorials/third-party/observability/

# HTTP接入
https://www.cloudwego.io/zh/docs/kitex/getting-started/tutorial/#%E6%9A%B4%E9%9C%B2-http-%E6%8E%A5%E5%8F%A3
当然也可以使用grpc gateway实现grpc http proxy，具体实现参考：feat/daheige/gateway分支
