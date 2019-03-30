// ApplicationProfilerMemory.go.

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

// Application Profiler for Memory.

type ApplicationProfilerMemory struct {
	ApplicationProfilerGeneric
}

func (this *ApplicationProfilerMemory) Configure(
	filePath string,
) error {

	var err error

	// Check.
	if len(filePath) == 0 {
		err = fmt.Errorf(
			ErrfFilePathNull,
			"ApplicationProfilerMemory.FilePath",
		)
		return err
	}

	// Configure.
	this.FilePath = filePath

	return nil
}

func (this *ApplicationProfilerMemory) Start() error {

	var err error

	// Check.
	if this.IsStarted {
		err = fmt.Errorf(ErrStarted)
		return err
	}

	// Start.
	this.File, err = os.Create(this.FilePath)
	if err != nil {
		return err
	}
	this.IsStarted = true

	return nil
}

func (this *ApplicationProfilerMemory) Use() error {

	var err error

	// Check.
	if !(this.IsStarted) {
		err = fmt.Errorf(ErrNotStarted)
		return err
	}

	// Use.
	runtime.GC()
	err = pprof.WriteHeapProfile(this.File)
	if err != nil {
		return err
	}

	return nil
}

func (this *ApplicationProfilerMemory) Stop() error {

	var err error

	// Check.
	if !(this.IsStarted) {
		err = fmt.Errorf(ErrNotStarted)
		return err
	}

	// Stop.
	err = this.File.Close()
	if err != nil {
		return err
	}
	this.IsStarted = false

	return nil
}
