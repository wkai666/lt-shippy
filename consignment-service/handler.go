package main

import (
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"log"
	pb "shippy/proto/consignment"
	vesselPb "shippy/proto/vessel"
)

// 定义微服务
type handler struct {
	session      *mgo.Session
	VesselClient vesselPb.VesselService
}

// GetRepo 从主会话中 clone 出新会话进行处理
// clone 新会话重用了主会话的 socket，避免了三次握手即资源的浪费
// copy 为会话创建新的 socket ，开销大
func (h *handler) GetRepo() IRepository {
	return &ConsignmentRepository{h.session.Clone()}
}

// CreateConsignment 托运新的货物
func (h *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	defer h.GetRepo().Close()
	// 检查是否有合适的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.GetWeight(),
	}

	vRes, err := h.VesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	// 货物被承运
	log.Printf("found vessle: %s\n", vRes.Vessel.Name)
	req.VesselId = vRes.Vessel.Id
	err = h.GetRepo().Create(req)
	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments 查看仓库中的货物
func (h *handler) GetConsignments(ctx context.Context, req *pb.Request, res *pb.Response) error {
	defer h.GetRepo().Close()
	consignments, err := h.GetRepo().GetAll()
	if err != nil {
		return err
	}

	res.Consignments = consignments
	return nil
}
