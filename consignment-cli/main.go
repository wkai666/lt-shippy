package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"log"
	"shippy/common"
	pb "shippy/proto/consignment"
)

func main() {
	// 客户端初始化
	service := micro.NewService()
	service.Init()
	client := pb.NewShippingService(common.ServiceConsignmentName, service.Client())

	// 装载货物
	consignment, err := load()
	if err != nil {
		log.Fatalf("load consignment err: %s\n", err)
	}

	// 客户端调用 rpc 存储货物到仓库
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment err: %v", err)
	}
	log.Printf("created: %v", resp.Created)

	// 调用 rpc 服务查看仓库中的货物
	resp, err = client.GetConsignments(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}

	for _, C := range resp.Consignments {
		log.Printf("%+v", C)
	}
}
