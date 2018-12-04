package main

import (
	vesselProto "github.com/heavenlwf/shippy/vessel-service/proto/vessel"
	pb "github.com/heavenlwf/shippy/consignment-service/proto/consignment"
	"gopkg.in/mgo.v2"
	"context"
	"log"
)

type service struct {
	session			*mgo.Session
	vesselClient	vesselProto.VesselServiceClient
}

func (s *service) GetRepo() Repository {
	return &ConsignmentRepository{s.session.Clone()}
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})

	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id

	err = repo.Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = req
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	repo := s.GetRepo()
	defer repo.Close()

	consignments, err := repo.GetAll()
	if err != nil {
		return nil
	}

	res.Consignments = consignments
	return nil
}