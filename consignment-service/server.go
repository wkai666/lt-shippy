package main

import (
	"github.com/asim/go-micro/v3"
	pb "shippy/consignment-service/proto/consignment"
	"shippy/consignment-service/service"
)

func main() {
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	server.Init()

	repo := service.Repository{}
	pb.RegisterShippingServiceHandler(server.Server(), &service.Service{repo})

	if err := server.Run(); err != nil {
		panic(err)
	}
}
