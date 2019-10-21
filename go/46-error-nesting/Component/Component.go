package c

import (
	"database/sql"
	"errors"

	e "github.com/legacy-vault/tests/go/46-error-nesting/Component/Error"
)

type Component struct {
}

func (this *Component) Init() (err error) {

	var internalError error

	// Error Simulation.
	internalError = sql.ErrNoRows
	if internalError != nil {
		err = e.New(
			e.StageInitialization,
			internalError,
		)
		return
	}

	return
}

func (this *Component) Fin() (err error) {

	var internalError error

	// Error Simulation.
	internalError = errors.New("The Computer has gone crazy")
	if internalError != nil {
		err = e.New(
			e.StageFinalization,
			internalError,
		)
		return
	}

	return
}
