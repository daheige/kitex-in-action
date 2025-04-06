FROM golang:1.22.12-alpine AS builder
LABEL author="daheige"

# 设置环境变量
ENV LANG="zh_CN.UTF-8"

# 配置环境变量（禁用 CGO、启用模块、设置代理）
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go mod download

# 编译构建应用程序
RUN cd /app/cmd/rpc && go build -ldflags="-w -s -extldflags '-static'" -o kitex-rpc

# 将上面构建好的二进制文件复制到容器中运行
#FROM alpine:latest
FROM alpine:3.21

# 设置时区和lang
ENV TZ=Asia/Shanghai LANG="zh_CN.UTF-8"

# 安装相关依赖
# 解决http x509证书问题，需要安装证书
# libc6-compat 安装DNS解析相关依赖（可选）
RUN echo "export LC_ALL=${LANG}"  >>  /etc/profile  \
    && apk update && apk upgrade  \
    && apk --no-cache add net-tools tzdata ca-certificates bash vim bash-completion curl  \
    && cp /usr/share/zoneinfo/${TZ} /etc/localtime  \
    && echo $TZ > /etc/timezone && apk del tzdata  \
    && rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $HOME/.cache

WORKDIR /app

# 服务端口
EXPOSE 8890 9093

# 将构建阶段的二进制文件复制到工作目录中
COPY --from=builder /app/cmd/rpc/kitex-rpc /app/kitex-rpc

ENTRYPOINT ["/app/kitex-rpc"]
