package service

import (
	"golang.org/x/net/context"
	pb "shippy/consignment-service/proto/consignment"
)

// Service 定义微服务
type Service struct {
	Repo Repository
}

// CreateConsignment 托运新的货物
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Created:     true,
		Consignment: consignment,
	}, nil
}

// GetConsignments 查看仓库中的货物
func (s *Service) GetConsignments(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	allConsignments := s.Repo.GetAll()
	return &pb.Response{
		Consignments: allConsignments,
	}, nil
}
