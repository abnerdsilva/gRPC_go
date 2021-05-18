package main

import (
	"context"
	"github.com/abnerdsilva/gRPC_go/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "hello gRPC",
	}

	res, err := client.RequestMessage(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetStatus())
}
