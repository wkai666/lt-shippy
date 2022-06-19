package main

import (
	"encoding/json"
	"io/ioutil"
	pb "shippy/proto/consignment"
)

const (
	DefaultInfoFile = "consignment.json"
)

func load() (*pb.Consignment, error) {
	// 准备货物
	data, err := ioutil.ReadFile(DefaultInfoFile)
	if err != nil {
		return nil, err
	}

	var consignment = &pb.Consignment{}
	if err = json.Unmarshal(data, consignment); err != nil {
		return nil, err
	}

	return consignment, nil
}
