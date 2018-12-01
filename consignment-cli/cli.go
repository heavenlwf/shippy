package main

import (
	pb "github.com/EwanValentine/shippy/consignment-service/proto/consignment"
	"io/ioutil"
	"log"
	"encoding/json"
	"os"
	"context"
	"github.com/micro/go-micro/cmd"
	microclient "github.com/micro/go-micro/client"
)

const (
	ADDRESS 			= "localhost:50051"
	DEFAULT_INFO_FILE 	= "consignment.json"
)

func parseConfig(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("read consignment.json Failed: %v", err)
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		log.Fatalf("json unmarshal failed: %v", err)
		return nil, err
	}

	return consignment, nil
}

func main() {
	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseConfig(infoFile)
	if err != nil {
		log.Fatalf("parseConfig failed: %v", err)
	}
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error: %v", err)
	}

	log.Printf("create consignment: %+v", resp.Consignment)

	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to get consignments: %v", err)
	}
	for _, c := range resp.Consignments {
		log.Printf("get %+v", c)
	}
}