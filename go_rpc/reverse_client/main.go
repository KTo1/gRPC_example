package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "go-rpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	var (
		hostName  = "localhost"
		pemFile   = "client.pem"
		keyFile   = "client.key"
		caFile    = "ca.crt"
		crtFolder = "cert"
	)

	cert, err := tls.LoadX509KeyPair(path.Join(crtFolder, pemFile), path.Join(crtFolder, keyFile))
	if err != nil {
		grpclog.Fatalf("Error loading certificate! %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(path.Join(crtFolder, caFile))
	if err != nil {
		grpclog.Fatalf("could not read CA certificate: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.DialOption{
		//grpc.WithInsecure(),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName:   hostName,
			Certificates: []tls.Certificate{cert},
			RootCAs:      certPool,
		})),
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
