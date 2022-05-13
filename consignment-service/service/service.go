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
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.Repo.Create(req)
	if err != nil {
		return err
	}

	res = &pb.Response{
		Created:     true,
		Consignment: consignment,
	}

	return nil
}

// GetConsignments 查看仓库中的货物
func (s *Service) GetConsignments(ctx context.Context, req *pb.Request, res *pb.Response) error {
	res = &pb.Response{
		Consignments: s.Repo.GetAll(),
	}

	return nil
}
