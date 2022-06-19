package main

import (
	"gopkg.in/mgo.v2"
	"shippy/common"
	pb "shippy/proto/consignment"
)

// IRepository 定义仓库
type IRepository interface {
	Create(consignment *pb.Consignment) error // 存放新货物
	GetAll() ([]*pb.Consignment, error)
	Close()
}

// ConsignmentRepository 存放多批货物，实现仓库接口
type ConsignmentRepository struct {
	session *mgo.Session
}

// Create 存放货物
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	return repo.collection().Insert(consignment)
}

// GetAll 获取所有货物
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var cons []*pb.Consignment
	err := repo.collection().Find(nil).All(&cons)
	return cons, err
}

// Close 关闭查询连接
func (repo *ConsignmentRepository) Close() {
	repo.session.Close()
}

// 建立查询链接
func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(common.DBName).C(common.DBCollectionConsignment)
}
