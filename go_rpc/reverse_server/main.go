package main

import (
	"fmt"
	pb "go-rpc/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	fmt.Print("server listening on 5300")

	pb.RegisterReverseServer(grpcServer, &server{})
	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
}

type server struct{}

func (s *server) DoSomething(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	message := request.Message

	response = &pb.Response{
		Message: message,
	}

	return response, nil
}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	n := 0

	randDelay(1500)

	// Create an array of runes to safely reverse a string.
	runeText := make([]rune, len(request.Message))

	for _, r := range request.Message {
		runeText[n] = r
		n++
	}

	// Reverse using runes.
	runeText = runeText[0:n]

	for i := 0; i < n/2; i++ {
		runeText[i], runeText[n-1-i] = runeText[n-1-i], runeText[i]
	}

	output := string(runeText)
	response = &pb.Response{
		Message: output,
	}

	return response, nil
}
