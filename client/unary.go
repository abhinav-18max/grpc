package main

import (
	"context"
	pb "github.com/abhinav-18max/grpc/proto"
	"log"
	"time"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	resp, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	log.Printf("Response: %s", resp.Message)

}
