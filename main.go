package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test/src"
)

func main() {
	//src.SetValue("my-cred", []byte("test-value"))
	//socket := "unix://test-socket.sock"
	//lis, err := net.Listen("unix", socket)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//s := src.Server{}
	//
	//grpcServer := grpc.NewServer()
	//proto.RegisterCredentialServiceServer(grpcServer, &s)
	//
	//fmt.Println("grpc server running on socket: ", socket)
	//
	//if err = grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//
	//}

	src.SetValue("my-cred", []byte("test-value"))
	if len(os.Args) == 1 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s <path-to-unix-socket>\n", os.Args[0])
		os.Exit(1)
	}

	s, err := src.New(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	s.Start()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	s.Stop()

}
