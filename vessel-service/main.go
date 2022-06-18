package main

import (
	"context"
	"errors"
	"github.com/asim/go-micro/v3"
	"log"
	pb "shippy/vessel-service/proto"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 选择最近一条容量，载重都符合的货轮
	for _, v := range repo.vessels {
		if v.Capacity >= spec.GetCapacity() && v.MaxWeight >= spec.GetMaxWeight() {
			return v, nil
		}
	}

	return nil, errors.New("No vessel.go can't be use")
}

// 定义货船服务
type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, res *pb.Response) error {
	// 调用内部方法查找
	v, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}

	res.Vessel = v
	return nil
}

const (
	ServiceNameVessel = "go.micro.srv.vessel"
)

func main() {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty Maboatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{
		vessels: vessels,
	}

	server := micro.NewService(
		micro.Name(ServiceNameVessel),
		micro.Version("latest"),
	)
	server.Init()

	// 将实现服务的 API 注册到服务端
	pb.RegisterVesselServiceHandler(server.Server(), &service{repo})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
