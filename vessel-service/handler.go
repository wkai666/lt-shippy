package main

import (
	"context"
	"gopkg.in/mgo.v2"
	pb "shippy/proto/vessel"
)

// 定义货船服务
type handler struct {
	session *mgo.Session
}

func (h *handler) GetRepo() Repository {
	return &VesselRepository{h.session.Clone()}
}

func (h *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	defer h.GetRepo().Close()

	if err := h.GetRepo().Create(req); err != nil {
		return err
	}

	res.Vessel = req
	res.Created = true

	return nil
}

func (h *handler) FindAvailable(ctx context.Context, spec *pb.Specification, res *pb.Response) error {
	// 调用内部方法查找
	v, err := h.GetRepo().FindAvailable(spec)
	if err != nil {
		return err
	}

	res.Vessel = v
	return nil
}
