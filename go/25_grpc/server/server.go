// server.go.

package main

// Built-In Imports.
import (
	"errors"
	"flag"
	"io"
	"log"
	"math"
	"net"
	"strconv"
)

// External Imports.
import "golang.org/x/net/context"
import "google.golang.org/grpc"

// Local Imports.
import pb "../protocol"

type Server struct{}

const SERVER_CONN_TYPE = "tcp"

// Gets the Square of a Number.
func (s *Server) GetSquare(
	ctx context.Context,
	data *pb.Data1) (*pb.Data2, error) {

	var x int32
	var y int64

	x = data.Value

	// Input Check is not needed.
	// (MaxInt32 * MaxInt32) is less than MaxInt64.

	// Calculate.
	y = int64(x)
	y = y * y

	return &pb.Data2{Value: y}, nil
}

// Gets the Sum of a Sequence of Numbers.
func (s *Server) GetSum(stream pb.DemoService_GetSumServer) error {

	var count int64
	var currentItem int32
	var currentItemLimit int64
	var data *pb.Data1
	var err error
	var result *pb.Data2
	var sum int64

	for {
		data, err = stream.Recv()
		if err != nil {
			if err == io.EOF {
				// End is reached. Send the Result.
				result = new(pb.Data2)
				result.Value = sum
				err = stream.SendAndClose(result)
				if err != nil {
					log.Println("Stream Send Error.", err)
					return err
				}
				return nil

			} else {
				// Read Error.
				log.Println("Stream Read Error.", err)
				return err
			}
		}
		// No Error. Process Data.
		currentItem = data.Value
		// Check for Overflow and sum up.
		currentItemLimit = math.MaxInt64 - sum
		if int64(currentItem) > currentItemLimit {
			// Overflow.
			err = errors.New("Overflow")
			return err
		}
		sum = sum + int64(currentItem)
		count++
	}
}

// Gets the Sequence of natural Numbers.
func (s *Server) ListNaturalNumbers(
	data *pb.Data3,
	stream pb.DemoService_ListNaturalNumbersServer) error {

	const MAX_COUNT = 100

	var count int64
	var err error
	var i int64
	var iMax int64
	var iMin int64
	var result *pb.Data2

	// Check Input Data.
	iMin = data.FirstValue
	iMax = data.LastValue
	if iMin <= 0 {
		err = errors.New("First Argument is not natural")
		return err
	}
	if iMax <= 0 {
		err = errors.New("Second Argument is not natural")
		return err
	}
	if iMin >= iMax {
		err = errors.New("Bad Arguments")
		return err
	}
	count = iMax - iMin + 1
	if count > MAX_COUNT {
		err = errors.New("The Task is too difficult for me")
		return err
	}

	// Iterate.
	result = new(pb.Data2)
	i = iMin
	for i <= iMax {
		result.Value = i
		err = stream.Send(result)
		if err != nil {
			log.Println("Stream Send Error.", err)
			return err
		}

		// Next i.
		i++
	}
	return nil
}

func main() {

	var address string
	var claSrvPort *int
	var err error
	var grpcServer *grpc.Server
	var listener net.Listener
	var server Server

	// Command Line Arguments.
	claSrvPort = flag.Int("port", 12345, "Server's TCP Port")
	flag.Parse()

	address = "0.0.0.0:" + strconv.FormatInt(int64(*claSrvPort), 10)
	listener, err = net.Listen(SERVER_CONN_TYPE, address)
	if err != nil {
		log.Fatal("Net Listen Error.", err)
	}

	grpcServer = grpc.NewServer()
	pb.RegisterDemoServiceServer(grpcServer, &server)
	grpcServer.Serve(listener)
}
