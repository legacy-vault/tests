// main.go.

package main

// Main.

import (
	"log"
)

func main() {

	var app Application
	var err error

	// Initialize Everything.
	err = mainInit(&app)
	if err != nil {
		log.Fatal(err)
	}

	// Work.
	work()

	// Use the Memory Profiler.
	err = app.Profilers.Memory.Use()
	if err != nil {
		log.Fatal(err)
	}

	// Finalize Everything.
	err = mainFin(&app)
	if err != nil {
		log.Fatal(err)
	}
}
