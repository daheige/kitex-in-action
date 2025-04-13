#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

proto_dir=$root_dir/idl

# 创建pb文件目录
pb_dir=${root_dir}/internal/pb
mkdir -p ${pb_dir}

# proto文件中的@inject_tag标签注入，主要用于json标签和
sh $root_dir/bin/protoc-inject-tag.sh

# gen request validator code
sh $root_dir/bin/validator-generate.sh
