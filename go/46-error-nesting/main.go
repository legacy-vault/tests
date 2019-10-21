package main

// Demonstration of nested Errors (Sub-Errors in Errors) in the 'Go' Language
// as of the Version 1.13.

import (
	"database/sql"
	"errors"
	"log"

	c "github.com/legacy-vault/tests/go/46/Component"
	e "github.com/legacy-vault/tests/go/46/Component/Error"
)

func main() {

	const NewLine = "\r\n"

	var component c.Component
	var err error

	// Initialization and its Error Check.
	err = component.Init()
	if err != nil {
		log.Println(
			// Get the full Text of an Error.
			err,
			NewLine,

			// Check the upper-level Error Type.
			errors.Is(err, e.ErrInitialization),
			errors.Is(err, e.ErrFinalization),
			errors.Is(err, e.ErrUnknown),
			NewLine,

			// Check that the internal Error is the SQL's 'No Rows' Error.
			errors.Is(err, sql.ErrNoRows),
			NewLine,

			// Get the internal Error to investigate the Situation.
			errors.Unwrap(err),
		)
	}
}

/*
	The Program outputs something similar to the following Text.
	---
	2019/10/21 23:00:00 Initialization: sql: no rows in result set
	 true false false
	 true
	 sql: no rows in result set
	---
*/
