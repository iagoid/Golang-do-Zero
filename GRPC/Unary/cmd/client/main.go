package main

import (
	"context"
	"log"
	"unary-gRPC/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Servidor do cliente, ele envia a requisição
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(conn)

	req := pb.HelloRequest{
		Name: "John Doe",
	}

	res, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
