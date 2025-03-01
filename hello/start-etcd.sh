# 本地运行etcd
container_name=etcd_test
container=$(docker ps -a | grep $container_name | awk '{print $1}')
if [ ${#container} -gt 0 ]; then
    docker rm -f $container_name
fi

# 运行etcd容器，端口12379
docker run -d \
  --name $container_name \
  --restart=always \
  -p 12379:2379 \
  -p 12380:2380 \
  quay.io/coreos/etcd:v3.5.1 \
  /usr/local/bin/etcd \
  --data-dir /etcd-data \
  --advertise-client-urls http://localhost:2379 \
  --listen-client-urls http://0.0.0.0:2379

echo "local etcd container run success!"
