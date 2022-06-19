package main

import (
	"github.com/asim/go-micro/v3"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"shippy/common"
	vesselPb "shippy/proto/vessel"
)

var (
	session      *mgo.Session
	vesselClient vesselPb.VesselService
)

func init() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = common.DBHost
	}

	// 连接 MongoDB
	session, err := CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session err: %s", err)
	}
	defer session.Close()

	// vessel 客户端初始化
	service := micro.NewService()
	service.Init()
	vesselClient = vesselPb.NewVesselService(common.ServiceVesselName, service.Client())
}
