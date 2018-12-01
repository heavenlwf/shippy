package main

import (
	pb "shippy/vessel-service/proto/vessel"
	"errors"
	"context"
	"log"
	"github.com/micro/go-micro"
	)

type Registory interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

// 接口实现
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, v := range repo.vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			return v, nil
		}
	}
	return nil, errors.New("No vessel can't be use")
}

// 定义货轮服务
type service struct {
	repo Registory
}

// 实现服务端
func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error  {
	// 调用内部方法查找
	v, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}

	resp.Vessel = v
	return nil
}
func main()  {
	// 停留在港口的货船，先写死
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels:vessels}
	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &service{repo:repo})
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
