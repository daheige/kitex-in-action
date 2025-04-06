#!/usr/bin/env bash
root_dir=$(cd "$(dirname "$0")"; cd ..; pwd)

image_name=kitex-develop
version=v1.0

container_name=kitex-dev-local
container=$(docker ps -a | grep $container_name | awk '{print $1}')
if [ ${#container} -gt 0 ]; then
    docker rm -f $container_name
fi

cd $root_dir
docker run --name=$container_name -itd $image_name:$version
echo "please exec cmd: docker exec -it $container_name /bin/bash enter into docker development"
