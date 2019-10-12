package main

import (
	pb "test/pb/2/pb"
)

func createStorageData() (data []*pb.Person) {

	var person *pb.Person

	// Person #1.
	person = &pb.Person{
		Id:   101,
		Name: "Alice",
		Phones: []*pb.Phone{
			&pb.Phone{
				Type:   pb.PhoneType(pb.PhoneType_CELLULAR),
				Number: "+7 999 111-11-11",
			},
			&pb.Phone{
				Type:   pb.PhoneType(pb.PhoneType_STATIONARY),
				Number: "+7 888 111-11-11",
			},
		},
	}
	data = append(data, person)

	// Person #2.
	person = &pb.Person{
		Id:   202,
		Name: "Bruce",
		Phones: []*pb.Phone{
			&pb.Phone{
				Type:   pb.PhoneType(pb.PhoneType_CELLULAR),
				Number: "+7 999 222-22-22",
			},
			&pb.Phone{
				Type:   pb.PhoneType(pb.PhoneType_STATIONARY),
				Number: "+7 888 222-22-22",
			},
		},
	}
	data = append(data, person)

	return
}
