package main

// Example of Context Usage in Go Language.

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Result struct {
	Error error
	Data  int
}

func main() {

	// Start the Generator.
	var outputChannel chan Result = make(chan Result, 1000)
	var ctx context.Context
	var cancelFunc context.CancelFunc
	ctx, cancelFunc = context.WithCancel(context.Background())
	defer cancelFunc()
	go generateData(ctx, outputChannel)

	// Let the Generator work for 3 Seconds, then stop it.
	time.Sleep(time.Millisecond * 3500)
	cancelFunc()

	// Get the Results.
	for item := range outputChannel {
		if item.Error != nil {
			log.Println(item.Error)
			break
		}
		fmt.Println(item.Data)
	}
}

// Generates some Data.
// The Results are sent into a Channel.
// Generator must be stopped using the Context.
func generateData(ctx context.Context, results chan Result) {

	defer close(results)

	var doneSignal bool
	var result Result
	for i := 1; i <= 10; i++ {

		// Simulate some hard Work.
		time.Sleep(time.Second * 1)

		// Should we continue?
		select {

		case _, doneSignal = <-ctx.Done():
			// Message from Context.
			if !doneSignal {
				// Shutdown.
				return
			} else {
				// Exception! This can not happen!
				var err error = fmt.Errorf("Context Error: %w", ctx.Err())
				results <- Result{Error: err}
				panic(err)
				return
			}

		default:
			// Normal Work.
			// Generate some Data, send the Results to the output Stream (Channel).
			result.Data = i
			results <- result
		}
	}
}
