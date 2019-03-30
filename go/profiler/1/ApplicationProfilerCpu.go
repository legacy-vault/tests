// ApplicationProfilerCpu.go.

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

// Application Profiler for CPU.

type ApplicationProfilerCpu struct {
	ApplicationProfilerGeneric
}

func (this *ApplicationProfilerCpu) Configure(
	filePath string,
) error {

	var err error

	// Check.
	if len(filePath) == 0 {
		err = fmt.Errorf(
			ErrfFilePathNull,
			"ApplicationProfilerCpu.FilePath",
		)
		return err
	}

	// Configure.
	this.FilePath = filePath

	return nil
}

func (this *ApplicationProfilerCpu) Start() error {

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
	err = pprof.StartCPUProfile(this.File)
	if err != nil {
		return err
	}

	return nil
}

func (this *ApplicationProfilerCpu) Stop() error {

	var err error

	// Check.
	if !(this.IsStarted) {
		err = fmt.Errorf(ErrNotStarted)
		return err
	}

	// Stop.
	pprof.StopCPUProfile()
	err = this.File.Close()
	if err != nil {
		return err
	}
	this.IsStarted = false

	return nil
}
