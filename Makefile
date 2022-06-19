.PHONY: pbuild help

pbuild:
	protoc --go_out=. --micro_out=. proto/consignment/consignment.proto
	protoc --go_out=. --micro_out=. proto/vessel/vessel.proto

help:
	@echo "make build 编译 proto 文件并编译 go代码生成二进制文件，最后编译 docker镜像"
	@echo "make pbuild 仅对 proto 文件进行编译"
	@echo "make gbuild 编译 go 代码生成二进制文件"
	@echo "make dbuild 根据 Dockerfile 编译相关镜像"