package service

import pb "shippy/consignment-service/proto/consignment"

// IRepository 定义仓库
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment
}

// Repository 存放多批货物，实现仓库接口
type Repository struct {
	consignments []*pb.Consignment
}

// Create 存放货物
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

// GetAll 获取所有货物
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}
