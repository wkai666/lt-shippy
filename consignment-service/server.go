package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "shippy/consignment-service/proto/consignment"
	"shippy/consignment-service/service"
)

const (
	PORT = ":50051"
)

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}
	log.Printf("listen on: %s\n", PORT)

	server := grpc.NewServer()
	repo := service.Repository{}

	// 向 grpc 服务器注册微服务，此时会将自己实现的 service 与 ShippingServiceServer 绑定
	pb.RegisterShippingServiceServer(server, &service.Service{repo})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
