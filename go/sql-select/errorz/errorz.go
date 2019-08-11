package errorz

import (
	"errors"
	"log"

	"test/sql-select/text"
)

const Separator = ";" + text.NewLine

// Stops the Application if an Error occurs.
func MustBeNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Combines two Errors into a single Error if they are both set.
func Combine(
	errorFirst error,
	errorSecond error,
) (err error) {

	if errorFirst == nil {
		return errorSecond
	}
	if errorSecond == nil {
		return errorFirst
	}
	return errors.New(errorFirst.Error() + Separator + errorSecond.Error())
}
