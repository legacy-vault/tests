package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/kr/pretty"
	"google.golang.org/grpc"

	pb "test/pb/2/pb"
)

func main() {

	var err error
	var grpcConnection *grpc.ClientConn
	var grpcClient pb.PersonServiceClient
	var grpcClientTargetAddress string
	var requestContext context.Context
	var requestCancelFunction context.CancelFunc
	var requestMessage *pb.GetPersonRequest
	var responseMessage *pb.GetPersonResponse

	// Create the gRPC Client.
	grpcClientTargetAddress = "localhost:3000"
	grpcConnection, err = grpc.Dial(grpcClientTargetAddress, grpc.WithInsecure())
	mustBeNoError(err)
	defer func() {
		var derr error
		derr = grpcConnection.Close()
		if derr != nil {
			log.Println(derr)
		}
	}()
	grpcClient = pb.NewPersonServiceClient(grpcConnection)

	// Send the Request and receive the Response.
	requestMessage = &pb.GetPersonRequest{
		PersonId: 101,
	}
	requestContext, requestCancelFunction = context.WithTimeout(
		context.Background(),
		time.Second*60,
	)
	defer requestCancelFunction()
	responseMessage, err = grpcClient.GetPerson(requestContext, requestMessage)
	mustBeNoError(err)
	log.Println(
		"Response:\r\n",
		pretty.Sprint(responseMessage),
	)
}

func mustBeNoError(
	err error,
) {
	if err != nil {
		log.Println("Error.", err.Error())
		os.Exit(1)
	}
}
