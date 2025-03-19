# kitex-in-action
    kitex实战项目开发
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
4. 安装grpc相关的go工具链
参考链接： https://www.cloudwego.io/zh/docs/kitex/tutorials/code-gen/code_generation/
```shell
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
# https://github.com/cloudwego/protoc-gen-validator
go install github.com/cloudwego/protoc-gen-validator@latest

# Reference: https://grpc.io/docs/languages/go/quickstart/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

# kitex实战demo
    见hello和kitex-example
