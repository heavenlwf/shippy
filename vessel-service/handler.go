package main

import (
	pb "github.com/heavenlwf/shippy/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"context"
)

type service struct {
	session *mgo.Session
}

func (s *service) GetRepo() Repository {
	return &VesselRepository{s.session.Clone()}
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vessel, err := repo.FindAvailable(req)
	if err != nil {
		return err
	}

	resp.Vessel = vessel
	return nil
}

func (s *service) Create(ctx context.Context, req *pb.Vessel, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	err := repo.Create(req)
	if err != nil {
		return err
	}

	resp.Vessel = req
	resp.Created = true
	return nil
}