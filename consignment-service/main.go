package main

import (
	"github.com/asim/go-micro/v3"
	"shippy/common"
	pb "shippy/proto/consignment"
)

func main() {
	server := micro.NewService(
		micro.Name(common.ServiceConsignmentName),
		micro.Version("latest"),
	)

	server.Init()
	pb.RegisterShippingServiceHandler(server.Server(), &handler{session, vesselClient})

	if err := server.Run(); err != nil {
		panic(err)
	}
}
