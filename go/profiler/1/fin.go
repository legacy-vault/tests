// fin.go.

package main

// Finalization.

func mainFin(
	app *Application,
) error {

	var err error

	// Stop the Profilers.
	err = app.Profilers.Cpu.Stop()
	if err != nil {
		return err
	}
	err = app.Profilers.Memory.Stop()
	if err != nil {
		return err
	}

	return nil
}
