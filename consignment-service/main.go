package main

import (
	vesselPb "github.com/heavenlwf/shippy/vessel-service/proto/vessel"
	pb "shippy/consignment-service/proto/consignment"
	"context"
	"log"
	"github.com/micro/go-micro"
)

// IRepository interface
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
}

type Repository struct {
	consignment []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error)  {
	repo.consignment = append(repo.consignment, consignment)
	log.Printf("Create consignment: %+v\n", consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() ([]*pb.Consignment, error)  {
	for _, v := range repo.consignment {
		log.Printf("Get All: %+v\n", v)
	}

	return repo.consignment, nil
}

type service struct {
	repo IRepository
	vesselClient vesselPb.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) error {
	// 检查是否有合适的货轮
	vReq := &vesselPb.Specification{
		Capacity: int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vResp, err := s.vesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		log.Fatalf("vesselClient FindAvailable err: %t", err)
		return err
	}
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)
	req.VesselId = vResp.Vessel.Id
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	resp.Created = true
	resp.Consignment = consignment
	//resp = &pb.Response{Created: true, Consignment: consignment}
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response)  error {
	allConsignments, err := s.repo.GetAll()
	if err != nil {
		return err
	}

	resp.Consignments = allConsignments
	return nil
}

func main() {
	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	server.Init()
	repo := &Repository{}
	vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessl", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &service{repo, vClient})
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
