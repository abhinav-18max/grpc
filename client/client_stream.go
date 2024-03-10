package main

import (
	"context"
	pb "github.com/abhinav-18max/grpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Couldn't connect to server: %v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Couldn't send message: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Streaming finished")
	if err != nil {
		log.Fatalf("Error while streaming: %v", err)
	}
	log.Printf("%v", res.Message)

}
