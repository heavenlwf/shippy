package main

import (
	pb "github.com/heavenlwf/shippy/vessel-service/proto/vessel"
	"log"
	"github.com/micro/go-micro"
	"os"
	)

const (
	DEFAULT_MONGO_HOST = "localhost:27017"
)

func createDummyData(repo Repository)  {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}

	for _, v := range vessels {
		repo.Create(v)
	}
}

func main()  {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = DEFAULT_MONGO_HOST
	}

	session, err := CreateSession(host)
	if err != nil {
		log.Fatalf("Connecting to datastore error: %v\n", err)
	}

	repo := &VesselRepository{session}
	createDummyData(repo)

	server := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)
	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &service{session})
	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
