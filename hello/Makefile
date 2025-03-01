gen-layout:
	# 仅仅是生成脚手架代码
	kitex -service greeter -module hello helloworld.proto

gen-pb:
	# 当proto文件发生更改，执行该命令，并实现对应的service方法即可
	kitex -module hello helloworld.proto

start-etcd:
	sh start-etcd.sh

enter-etcd:
	docker exec -it etcd_test /bin/bash

.PHONY: gen-layout gen-pb
