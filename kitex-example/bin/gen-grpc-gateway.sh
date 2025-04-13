#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)
protoExec=$(which "protoc")

proto_dir=$root_dir/idl

# 创建pb文件目录
pb_dir=${root_dir}/internal/pb/gateway
mkdir -p ${pb_dir}

if [ -z $protoExec ]; then
    echo 'Please install protoc'
    echo "Please look kitex-in-action/readme.md to install protoc"
    echo "if you use centos7,please look https://github.com/daheige/go-proj/blob/master/docs/centos7-protoc-install.md"
    exit 0
fi

$protoExec -I $proto_dir \
    --go_out $pb_dir --go_opt paths=source_relative \
    --go-grpc_out $pb_dir --go-grpc_opt paths=source_relative \
    $proto_dir/*.proto

$protoExec -I $proto_dir --grpc-gateway_out $pb_dir \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    $proto_dir/*.proto

go mod tidy

echo "\n\033[0;32mGenerating grpc gateway codes successfully!\033[39;49;0m\n"
exit 0