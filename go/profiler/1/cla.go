// cla.go.

package main

// Command Line Arguments.

import (
	"flag"
	"fmt"
)

const (
	Cla_ProfilerCpuFilePath       = "profiler_cpu"
	Cla_ProfilerMemoryCpuFilePath = "profiler_mem"
)

func claInit(
	app *Application,
) error {

	var cpuProfilerFilePath *string
	var err error
	var memoryProfilerFilePath *string

	// Input Parameters Check.
	if app == nil {
		err = fmt.Errorf(
			ErrfPointerNull,
			"app",
		)
		return err
	}

	// Command Line Arguments parsing.
	cpuProfilerFilePath = flag.String(
		Cla_ProfilerCpuFilePath,
		"",
		"CPU Profiler Output File",
	)
	memoryProfilerFilePath = flag.String(
		Cla_ProfilerMemoryCpuFilePath,
		"",
		"Memory Profiler Output File",
	)
	flag.Parse()

	// Configure the CPU Profiler.
	if cpuProfilerFilePath == nil {
		err = fmt.Errorf(
			ErrfPointerNull,
			"cpuProfilerFilePath",
		)
		return err
	}
	err = app.Profilers.Cpu.Configure(*cpuProfilerFilePath)
	if err != nil {
		return err
	}

	// Configure the Memory Profiler.
	if memoryProfilerFilePath == nil {
		err = fmt.Errorf(
			ErrfPointerNull,
			"memoryProfilerFilePath",
		)
		return err
	}
	err = app.Profilers.Memory.Configure(*memoryProfilerFilePath)
	if err != nil {
		return err
	}

	return nil
}
