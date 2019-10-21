package e

import (
	"errors"
)

const (
	// 'Error Stage' Enumeration Simulation.
	// Unfortunately, 'Go' Language does not know about Enumerations.
	StageInitialization = 1
	StageFinalization   = 2

	ErrTextInitialization = "Initialization"
	ErrTextFinalization   = "Finalization"
	ErrTextUnknown        = "Unknown Error"

	ErrTextInternalSeparator = ": "
)

type Error struct {
	// Stage at which an Error has occurred.
	stage byte

	// An internal Error.
	internalError error
}

// Unfortunately, the 'Go' Language can not make Variables constant.
// This is very stupid...
var (
	ErrInitialization = errors.New(ErrTextInitialization)
	ErrFinalization   = errors.New(ErrTextFinalization)
	ErrUnknown        = errors.New(ErrTextUnknown)
)

// A simple Error Constructor.
func New(
	stage byte,
	internalError error,
) Error {

	return Error{
		stage:         stage,
		internalError: internalError,
	}
}

// Implementation of the 'Error' Method of the built-in 'error' Interface.
// Returns the full Error Text including the internal Error.
func (this Error) Error() (errorText string) {

	switch this.stage {

	case StageInitialization:
		errorText = ErrTextInitialization + ErrTextInternalSeparator +
			this.internalError.Error()
		return

	case StageFinalization:
		errorText = ErrTextFinalization + ErrTextInternalSeparator +
			this.internalError.Error()
		return

	default:
		return ErrTextUnknown
	}
}

// Returns the short Error Text, without the internal Error.
func (this Error) ShortError() (errorText string) {

	switch this.stage {

	case StageInitialization:
		return ErrTextInitialization

	case StageFinalization:
		return ErrTextFinalization

	default:
		return ErrTextUnknown
	}
}

// Implementation of the 'Is' Method used by the built-in 'errors' Package.
func (this Error) Is(
	sample error,
) bool {

	//sampleError, ok := sample.(Error)
	//if !ok {
	//	return false
	//}
	//
	//return (sampleError.Error() == e.Error())

	return (this.ShortError() == sample.Error())
}

// Implementation of the 'Unwrap' Method used by the built-in 'errors' Package.
func (this Error) Unwrap() error {
	return this.internalError
}
