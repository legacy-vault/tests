// main.go.

// Block CRC Test.

// Date: 2018-09-22.

package main

import (
	"fmt"
	"os"
)

// Default Configuration Values (before Configuration has been read).
const Default_VerboseMode = false

var verboseMode bool

// Program's Entry Point.
func main() {

	var err error

	// Set Verbose Mode before it has been set by Comman Line Arguments.
	verboseMode = Default_VerboseMode

	// Initialize Application Stages.
	err = stagesInit()
	if err != nil {
		panic(err)
	}

	// Initialize Everything.
	runStage(Stage_Initialization)
	defer FinDeferred()

	// Read O.S. Command Line Arguments.
	runStage(Stage_CommandLineArguments)

	// Run Action.
	runStage(actionStage)
}

// Initializations.
func Init() error {

	var err error

	// Self-Check.
	err = selfCheck()
	if err != nil {
		return err
	}

	return nil
}

// Finalizations.
func Fin() error {

	var err error

	// Do Something.
	err = nil // A Plug.
	if err != nil {
		return err
	}

	os.Exit(0)

	return nil
}

// Deferred Finalizations.
func FinDeferred() {

	runStage(Stage_Finalization)
}

// Does Nothing.
func void() error {

	return nil
}

// Runs the Stage.
func runStage(stage int) {

	var err error
	var f StageFunction
	var msg string

	// Check Stage Index.
	if (stage < Stage_Index_Min) || (stage > Stage_Index_Max) {
		stage = Stage_Unknown
	}

	// Get Stage Function.
	f = stageFucntions[stage]

	// Run Stage Function.
	if verboseMode {
		//msg = stageStatusText(stage, Status_Started)
		//fmt.Println(msg)
	}
	err = f()
	if err != nil {
		if verboseMode {
			msg = stageStatusText(stage, Status_Failed)
			fmt.Println(msg)
		}
		panic(err)
	}
	if verboseMode {
		msg = stageStatusText(stage, Status_Finished)
		fmt.Println(msg)
	}
}
