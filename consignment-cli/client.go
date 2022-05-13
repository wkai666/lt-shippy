package main

import (
	"context"
	"encoding/json"
	"github.com/asim/go-micro/v3"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	ServiceName     = "go.micro.srv.consignment"
	DefaultInfoFile = "consignment.json"
)

func main() {
	// 准备货物
	infoFile := DefaultInfoFile
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}
	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file err: %v", err)
	}

	// 客户端初始化
	service := micro.NewService()
	service.Init()
	client := pb.NewShippingService(ServiceName, service.Client())

	// 客户端调用 rpc 存储货物到仓库
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment err: %v", err)
	}

	log.Printf("created: %t", resp.Created)

	// 调用 rpc 服务查看仓库中的货物
	resp, err = client.GetConsignments(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("failed to list consignments: %v", err)
	}

	for _, C := range resp.Consignments {
		log.Printf("%+v", C)
	}
}

func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	if err := json.Unmarshal(data, &consignment); err != nil {
		return nil, err
	}

	return consignment, nil
}
