package main

import (
	"github.com/asim/go-micro/v3/util/log"
	"gopkg.in/mgo.v2"
	"os"
	"shippy/common"
)

var (
	session *mgo.Session
)

func init() {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = common.DBHost
	}

	session, err := common.CreateSession(dbHost)
	if err != nil {
		log.Fatalf("create session err: %s\n", err)
	}
	defer session.Close()
}
