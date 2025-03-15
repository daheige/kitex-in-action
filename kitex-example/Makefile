# go.mod文件中的模块名字
module_name := kitex-example
# 服务名字对应proto文件的服务名的小写名字
service_name := greeter
# proto文件名字
proto_name := helloworld.proto

# 仅仅是生成脚手架代码
gen-layout:
	kitex -service ${service_name} -module ${module_name} -I idl idl/${proto_name}
	sh bin/gen-layout.sh ${module_name}

# 当proto文件发生更改，执行该命令，并实现对应的service方法即可
gen-pb:
	kitex -module ${module_name} -I idl idl/${proto_name}
	sh bin/gen-pb.sh ${module_name}
