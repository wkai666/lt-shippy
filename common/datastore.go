package common

import (
	"gopkg.in/mgo.v2"
	"log"
)

func CreateSession(host string) (*mgo.Session, error) {
	s, err := mgo.Dial(host)
	if err != nil {
		log.Fatalf("mongo dail error: %s", err)
	}
	s.SetMode(mgo.Monotonic, true)
	return s, nil
}
