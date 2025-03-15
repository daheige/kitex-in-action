#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=kitex-rpc
version=v1.0
cd $root_dir

container_name=kitex-rpc-svc
container=$(docker ps -a | grep $container_name | awk '{print $1}')
if [ ${#container} -gt 0 ]; then
    docker rm -f $container_name
fi

docker run --name=$container_name -p 8890:8890 -itd $image_name:$version