package main

// A simple Example of a simple gRPC Server.

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "test/pb/2/pb"
)

const (
	RpcServerAddress = "0.0.0.0:3000"
)

func main() {

	var err error
	var grpcServer *grpc.Server
	var grpcPersonServer *GrpcPersonServer
	var listener net.Listener

	// Create a simple gRPC person Server.
	grpcServer = grpc.NewServer()
	grpcPersonServer = &GrpcPersonServer{
		storage: createStorageData(),
	}
	pb.RegisterPersonServiceServer(grpcServer, grpcPersonServer)

	// Start the gRPC Server.
	listener, err = net.Listen("tcp", RpcServerAddress)
	mustBeNoError(err)
	err = grpcServer.Serve(listener)
	mustBeNoError(err)
}

func mustBeNoError(
	err error,
) {
	if err != nil {
		log.Println("Error.", err.Error())
		os.Exit(1)
	}
}
