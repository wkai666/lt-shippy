package main

import (
	"github.com/asim/go-micro/v3"
	"log"
	"shippy/common"
	pb "shippy/proto/vessel"
)

func main() {
	server := micro.NewService(
		micro.Name(common.ServiceVesselName),
		micro.Version("latest"),
	)
	server.Init()

	// 将实现服务的 API 注册到服务端
	pb.RegisterVesselServiceHandler(server.Server(), &handler{session})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
