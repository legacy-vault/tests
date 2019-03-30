// init.go.

package main

// Initialization.

func mainInit(
	app *Application,
) error {

	var err error

	// Initialize Command Line Arguments.
	err = claInit(app)
	if err != nil {
		return err
	}

	// Start the Profilers.
	err = app.Profilers.Cpu.Start()
	if err != nil {
		return err
	}
	err = app.Profilers.Memory.Start()
	if err != nil {
		return err
	}

	return nil
}
