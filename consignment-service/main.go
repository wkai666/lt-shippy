package main

import (
	"github.com/asim/go-micro/v3"
	"shippy/consignment-service/handler"
	pb "shippy/consignment-service/proto/consignment"
	vesselPb "shippy/vessel-service/proto"
)

const (
	ServiceVesselName      = "go.micro.srv.vessel"
	ServiceConsignmentName = "go.micro.srv.consignment"
)

var (
	repo         handler.Repository
	vesselClient vesselPb.VesselService
)

func init() {
	repo = handler.Repository{}

	// vessel 客户端初始化
	service := micro.NewService()
	service.Init()
	vesselClient = vesselPb.NewVesselService(ServiceVesselName, service.Client())
}

func main() {
	server := micro.NewService(
		micro.Name(ServiceConsignmentName),
		micro.Version("latest"),
	)

	server.Init()
	pb.RegisterShippingServiceHandler(server.Server(), &handler.Service{repo, vesselClient})

	if err := server.Run(); err != nil {
		panic(err)
	}
}
