package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/kr/pretty"

	pbPerson "../pb"
)

func main() {

	var err error
	var personA *pbPerson.Person
	var personAEncoded []byte
	var personB *pbPerson.Person

	// Prepare a sample Object.
	personA = new(pbPerson.Person)
	personA.Id = 123
	personA.Name = "John"
	personA.Phones = make([]*pbPerson.Phone, 0)
	personA.Phones = append(personA.Phones,
		&pbPerson.Phone{
			Type:   pbPerson.PhoneType(pbPerson.PhoneType_CELLULAR),
			Number: "+7 916 000-00-00",
		},
		&pbPerson.Phone{
			Type:   pbPerson.PhoneType(pbPerson.PhoneType_STATIONARY),
			Number: "+7 111 888-00-00",
		},
	)

	// Encode the Object.
	personAEncoded, err = proto.Marshal(personA)
	mustBeNoError(err)

	// Decode the Object Data.
	personB = new(pbPerson.Person)
	err = proto.Unmarshal(personAEncoded, personB)
	mustBeNoError(err)

	// Verification.
	// Actually, two Objects using the Google's "Protocol Buffers" are not
	// exactly the same internally. Unfortunately...
	mustBeEqual(personA, personB)

	// Exit.
	os.Exit(0)
}

func mustBeNoError(
	err error,
) {
	if err != nil {
		log.Println("Error.", err.Error())
		os.Exit(1)
	}
}

func mustBeEqual(
	x interface{},
	y interface{},
) {
	if !reflect.DeepEqual(x, y) {
		log.Println(
			"Values are different.\r\n" +
				fmt.Sprintf(
					"X: %+v, \r\nY: %+v, \r\nDifference: %+v.",
					pretty.Sprint(x),
					pretty.Sprint(y),
					pretty.Diff(x, y),
				),
		)
		os.Exit(1)
	}
}
