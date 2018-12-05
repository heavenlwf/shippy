package main

import (
	vesselPb "github.com/heavenlwf/shippy/vessel-service/proto/vessel"
	pb "github.com/heavenlwf/shippy/consignment-service/proto/consignment"
		"log"
	"github.com/micro/go-micro"
	"os"
	)

const (
	DEFAULT_MONGO_HOST = "localhost:27017"
)


// IRepository interface
type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
}

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = DEFAULT_MONGO_HOST
	}
	session, err := CreateSession(dbHost)
	defer session.Close()
	if err != nil {
		log.Fatalf("create session error: %v\n", err)
	}

	server := micro.NewService(
		// 必须和 consignment.proto 中的 package 一致
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	// 解析命令行参数
	server.Init()

	vClient := vesselPb.NewVesselServiceClient("go.micro.srv.vessel", server.Client())
	pb.RegisterShippingServiceHandler(server.Server(), &service{session, vClient})

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
