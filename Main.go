package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"test/proto"
	"test/src"
)

func main(){
	src.SetValue("my-cred", "test-value")
	port:= 10000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := src.Server{}

	grpcServer := grpc.NewServer()
	proto.RegisterCredentialServiceServer(grpcServer, &s)

	fmt.Println("grpc server running on port: ", port)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)

	}
}
