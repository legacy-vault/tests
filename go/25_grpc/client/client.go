// client.go.

package main

// Built-In Imports.
import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"
)
// External Imports.
import "golang.org/x/net/context"
import "google.golang.org/grpc"

// Local Imports.
import pb "../protocol"

func main() {

	var address string
	var cancelFunc context.CancelFunc
	var client pb.DemoServiceClient
	var connection *grpc.ClientConn
	var ctx context.Context
	var err error
	var i int
	var param1 pb.Data1
	var param2 pb.Data3
	var rcvStream pb.DemoService_ListNaturalNumbersClient
	var result *pb.Data2
	var rndGen *rand.Rand
	var sendStream pb.DemoService_GetSumClient
	var task3Values []int32

	// Connect.
	address = "127.0.0.1:12345"
	connection, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("Connection Error.", err)
	}
	defer connection.Close()
	client = pb.NewDemoServiceClient(connection)

	// Context.
	ctx, cancelFunc = context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	// 1. Call a simple remote Procedure.
	fmt.Print("Call #1. ")
	param1.Value = 7
	fmt.Printf("Square of %v: ", param1.Value)
	result, err = client.GetSquare(ctx, &param1)
	if err != nil {
		log.Fatal("RPC Error.", err)
	}
	fmt.Printf("Result: %v.\r\n", result.Value)

	// 2.1. Call a remote Procedure which returns a Stream. Normal Parameters.
	fmt.Print("Call #2.1. ")
	param2.FirstValue = 3
	param2.LastValue = 7
	rcvStream, err = client.ListNaturalNumbers(context.Background(), &param2)
	if err != nil {
		log.Fatal("RPC Stream Error.", err)
	}
	fmt.Printf("Natural Numbers from %v to %v:",
		param2.FirstValue, param2.LastValue)
	for {
		result, err = rcvStream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("RPC Error. %v", err)
				break
			}
		}
		fmt.Printf(" %v", result.Value)
	}
	fmt.Printf(".\r\n")

	// 2.2. Call a remote Procedure which returns a Stream. Bad Parameters.
	fmt.Print("Call #2.2. ")
	param2.FirstValue = -6
	param2.LastValue = 7
	rcvStream, err = client.ListNaturalNumbers(context.Background(), &param2)
	if err != nil {
		log.Fatal("RPC Stream Error.", err)
	}
	fmt.Printf("Natural Numbers from %v to %v:",
		param2.FirstValue, param2.LastValue)
	for {
		result, err = rcvStream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("RPC Error. %v", err)
				break
			}
		}
		fmt.Printf(" %v", result.Value)
	}
	fmt.Printf(".\r\n")

	// 3. Call a remote Procedure which receives a Stream.
	fmt.Print("Call #3. Sending Data...")
	task3Values = make([]int32, 5)
	rndGen = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i = 0; i < 5; i++ {
		task3Values[i] = rndGen.Int31()
		fmt.Printf(" %v", task3Values[i])
	}
	fmt.Printf(". ")
	sendStream, err = client.GetSum(context.Background())
	if err != nil {
		log.Fatal("Stream Send Error.", err)
	}
	for i = 0; i < 5; i++ {
		param1.Value = task3Values[i]
		err = sendStream.Send(&param1)
		if err != nil {
			log.Fatal("Stream Send Error.", err)
		}
	}
	result, err = sendStream.CloseAndRecv()
	if err != nil {
		log.Fatal("Stream Read Error.", err)
	}
	fmt.Printf("Result: %v.\r\n", result.Value)
}
