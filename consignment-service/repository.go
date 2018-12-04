package main

import (
	pb "github.com/heavenlwf/shippy/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
)

const (
	DB_NAME 		= "shippy"
	CON_COLLECTION 	= "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	session *mgo.Session
}

// 接口实现
func (repo *ConsignmentRepository) Create(c *pb.Consignment) error {
	return repo.collection().Insert(c)
}

func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	err := repo.collection().Find(nil).All(&consignments)
	return consignments, err
}

func (repo *ConsignmentRepository) Close()  {
	repo.session.Close()
}

func (repo *ConsignmentRepository) collection() *mgo.Collection {
	return repo.session.DB(DB_NAME).C(CON_COLLECTION)
}