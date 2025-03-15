#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)
module_name=$1
if [ -z $module_name ]; then
   echo "please input module_name"
   exit 0
fi

mkdir -p ${root_dir}/internal/interfaces/rpc
mkdir -p ${root_dir}/cmd/rpc
mkdir -p ${root_dir}/internal/application
mkdir -p ${root_dir}/internal/domain
mkdir -p ${root_dir}/internal/domain/entity
mkdir -p ${root_dir}/internal/domain/repo
mkdir -p ${root_dir}/internal/infras
mkdir -p ${root_dir}/internal/interfaces/rpc

# 创建pb文件目录
pb_dir=${root_dir}/internal/pb
mkdir -p ${pb_dir}

cp -R ${root_dir}/kitex_gen/pb/* ${pb_dir}
rm -rf ${root_dir}/kitex_gen

echo "fix the package path generated automatically by proto file code"
os=`uname -s`
if [ $os == "Darwin" ];then
    # mac os LC_CTYPE config
    export LC_CTYPE=C

    # mac os
    sed -i "" "s/${module_name}\/kitex_gen\/pb/${module_name}\/internal\/pb/g" `grep ${module_name}/kitex_gen/pb -rl ${pb_dir}`
else
    sed -i "s/${module_name}\/kitex_gen\/pb/${module_name}\/internal\/pb/g" `grep ${module_name}/kitex_gen/pb -rl ${pb_dir}`
fi

cp ${root_dir}/bin/main_code.tpl ${root_dir}/cmd/rpc/main.go
cp ${root_dir}/bin/greeter_impl.tpl ${root_dir}/internal/interfaces/rpc/greeter_service.go
rm -rf ${root_dir}/handler.go
rm -rf ${root_dir}/script
rm -rf ${root_dir}/build.sh
rm -rf ${root_dir}/main.go

# 获取对应的go package
go mod tidy

echo "\n\033[0;32mGenerating layout codes successfully!\033[39;49;0m\n"
exit 0
