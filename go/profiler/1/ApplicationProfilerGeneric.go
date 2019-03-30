// ApplicationProfilerGeneric.go.

package main

import "os"

// Application Profiler (Generic).

const ErrNotStarted = "Not Started"
const ErrStarted = "Already Started"

type ApplicationProfilerGeneric struct {
	File      *os.File
	FilePath  string
	IsStarted bool
}
