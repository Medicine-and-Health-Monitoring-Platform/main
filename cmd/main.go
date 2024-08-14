package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/config"
	pb "main/genproto/health_analytics"
	"main/kafka/consumer"
	"main/pkg/logger"
	"main/service"
	"main/storage"
	"main/storage/mongosh"
	"net"

	"google.golang.org/grpc"
)

var Mdb storage.IStorage
var err error

func main() {
	Mdb, err = mongosh.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer Mdb.Clouse()

	conf := config.Load()
	fmt.Println("Starting server ...")

	lis, err := net.Listen("tcp", conf.HTTP_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()
	logger := logger.NewLogger()

	md := service.NewHealthService(logger, Mdb)
	server := grpc.NewServer()
	pb.RegisterHealthAnalyticsServiceServer(server, md)

	reader, err := consumer.NewKafkaConsumInit([]string{"kafka:9092"}, "create", "group")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	go func() {
		reader.ComsumeMessages(ComsumeMessage)
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func ComsumeMessage(message []byte) {
	fmt.Println("Consume message: ", string(message))

	var req pb.AddMedicalRecordRequest
	err = json.Unmarshal(message, &req)
	if err != nil {
		fmt.Printf("error unmarshalling message: %v\n", err)
		return
	}

	res,err := Mdb.MedicalRecords().AddMedicalRecord(context.Background(), &req)
	if err != nil {
		fmt.Printf("error delete language: %v\n", err)
		return
	}

	fmt.Println("Medical record added successfully", "record_id", res.GetRecordId())


}
