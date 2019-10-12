package main

import (
	"context"
	"log"

	pb "test/pb/2/pb"
)

type GrpcPersonServer struct {
	pb.PersonServiceServer
	storage []*pb.Person
}

func (rs *GrpcPersonServer) GetPerson(
	ctx context.Context,
	request *pb.GetPersonRequest,
) (*pb.GetPersonResponse, error) {

	var person *pb.Person
	var personIsFound bool

	log.Printf("GetPerson(%v).", request.PersonId) //!

	// Find a Person in the Storage.
	for _, person = range rs.storage {
		if person.Id == request.PersonId {
			personIsFound = true
			break
		}
	}

	// No Person is found.
	if !personIsFound {
		return &pb.GetPersonResponse{
			Persons: []*pb.Person{},
		}, nil
	}

	// A Person is found.
	return &pb.GetPersonResponse{
		Persons: []*pb.Person{
			person,
		},
	}, nil
}
