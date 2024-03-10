package main

import (
	"context"
	pb "github.com/abhinav-18max/grpc/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Couldn't connect to server: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error getting message: %v", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming finished")
}
