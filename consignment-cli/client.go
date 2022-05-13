package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	pb "shippy/consignment-service/proto/consignment"
)

const (
	Address         = "localhost:50051"
	DefaultInfoFile = "consignment.json"
)

func main() {
	// 连接到 grpc服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect err: %v", err)
	}
	defer conn.Close()

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
	client := pb.NewShippingServiceClient(conn)

	// 客户端调用 rpc 存储货物到仓库
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment err: %v", err)
	}
	// 查看货物是否托运成功
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
