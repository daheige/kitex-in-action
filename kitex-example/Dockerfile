FROM golang:1.22.12-bullseye as builder
LABEL authors="daheige"

# 设置环境变量
ENV LANG="zh_CN.UTF-8"

# 设置环境变量
ENV CGO_ENABLED=0 GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

# 编译构建应用程序
RUN cd /app/cmd/rpc && go build -ldflags="-s -w" -o kitex-rpc

# 将上面构建好的二进制文件复制到容器中运行
FROM debian:bullseye-slim

# 设置时区和lang
ENV TZ=Asia/Shanghai LANG="zh_CN.UTF-8"

WORKDIR /app

# 服务端口
EXPOSE 8890

# 设置deb镜像源，这里我使用aliyun的镜像
RUN echo "deb http://mirrors.aliyun.com/debian bullseye main" > /etc/apt/sources.list &&  \
    echo "deb http://mirrors.aliyun.com/debian-security bullseye-security main" >> /etc/apt/sources.list &&  \
    echo "deb http://mirrors.aliyun.com/debian bullseye-updates main" >> /etc/apt/sources.list &&  \
    apt-get update && apt-get install -y ca-certificates vim bash curl net-tools wget \
    apt-transport-https && update-ca-certificates && apt-get clean &&  \
    rm -rf /var/lib/apt/lists/* && mkdir -p /app/config

# 将构建阶段的二进制文件复制到工作目录中
COPY --from=builder /app/cmd/rpc/kitex-rpc /app/kitex-rpc

ENTRYPOINT ["/app/kitex-rpc"]
