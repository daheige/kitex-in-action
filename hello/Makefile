# 仅仅是生成脚手架代码
gen-layout:
	kitex -service greeter -module hello -I idl idl/helloworld.proto

# 当proto文件发生更改，执行该命令，并实现对应的service方法即可
gen-pb:
	kitex -module hello -I idl idl/helloworld.proto

start-etcd:
	sh start-etcd.sh

enter-etcd:
	docker exec -it etcd_test /bin/bash

build:
	sh build.sh

# 先构建，再运行
run: build
	sh output/bootstrap.sh

.PHONY: gen-layout gen-pb build
