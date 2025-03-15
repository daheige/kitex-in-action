#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)
module_name=$1
if [ -z $module_name ]; then
   echo "please input module_name"
   exit 0
fi

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

# 获取对应的go package
go mod tidy

echo "\n\033[0;32mGenerating pb codes successfully!\033[39;49;0m\n"
exit 0
