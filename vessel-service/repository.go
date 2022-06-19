package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"shippy/common"
	pb "shippy/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
	Close()
}

type VesselRepository struct {
	session *mgo.Session
}

func (repo *VesselRepository) Create(v *pb.Vessel) error {
	return repo.collection().Insert(v)
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	if err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.GetCapacity()},
		"maxweight": bson.M{"$gte": spec.GetMaxWeight()},
	}).One(&vessel); err != nil {
		return nil, err
	}

	return vessel, nil
}

func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(common.DBName).C(common.DBCollectionVessel)
}
