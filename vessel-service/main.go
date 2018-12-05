package main

import (
	pb "shippy/vessel-service/proto/vessel"
	"log"
	"github.com/micro/go-micro"
	)

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
