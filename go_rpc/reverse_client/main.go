package main

import (
	"context"
	"fmt"
	pb "go-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"os"
	"path"
)

func main() {
	var (
		hostName  = "localhost"
		pemFile   = "server.pem"
		crtFolder = "cert"
	)

	cred, err := credentials.NewClientTLSFromFile(path.Join(crtFolder, pemFile), hostName)
	if err != nil {
		grpclog.Fatalf("Error loading certificate! %v", err)
	}

	opts := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithTransportCredentials(cred),
	}

	args := os.Args
	conn, err := grpc.Dial("127.0.0.1:5300", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewReverseClient(conn)
	request := &pb.Request{
		Message: args[1],
	}
	response, err := client.Do(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response.Message)
}
