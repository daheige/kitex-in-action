# kitex-in-action
为了更好地帮助小伙伴们使用kitex，降低学习成本，以及减少接入成本（最终的目标：让开发人员能更好地聚焦业务逻辑开发）。<br/>
因此，我将kitex在使用过程中的各个方面实战进行了开源，主要包含以下内容：
- kitex工具使用
- 如何快速编写微服务接口
- 如何编写微服务server端中间件
- 如何快速接入日志服务
- 如何接入服务监控和metrics数据指标上报以及pprof接入
- 如何快速接入grpc http gateway
- 如何编写客户端代码调用微服务接口

# kitex参考文档
https://www.cloudwego.io/zh/docs/kitex/

# 环境准备
1. 安装go
进入 https://go.dev/dl/ 官方网站，根据系统安装不同的go版本，这里推荐在linux或mac系统go
2. 设置GOPROXY
```shell
go env -w GOPROXY=https://goproxy.cn,direct
```
3. 安装kitex工具
```shell
# go install github.com/cloudwego/kitex/tool/cmd/kitex@v0.13.1
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
```
3. 安装protoc工具
- mac系统安装方式如下：
```shell
brew install protobuf
```
- linux系统安装方式如下：
```shell
# Reference: https://grpc.io/docs/protoc-installation/
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
unzip -o protoc-3.15.8-linux-x86_64.zip -d $HOME/.local
export PATH=~/.local/bin:$PATH # Add this to your `~/.bashrc`.
protoc --version
libprotoc 3.15.8
```

4. 安装kitex和grpc相关的go工具链
参考链接： https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/code_generation/
```shell
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
# https://github.com/cloudwego/protoc-gen-validator
go install github.com/cloudwego/protoc-gen-validator@latest

# go gRPC tools
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
# Reference: https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

# kitex实战demo
    见hello和kitex-example
