// self_check.go.

// Program's Self Check

// Date: 2018-09-22.

package main

import "errors"

// Error Messages.
const ErrMsgSelfCheckStageIndices = "Stage Indices Self Check Error"
const ErrMsgSelfCheckStatusIndices = "Stage Status Indices Self Check Error"

// Checks Application's Internals.
func selfCheck() error {

	var err error

	// Check Stage Indices.
	err = checkStageIndices()
	if err != nil {
		return err
	}

	return nil
}

// Checks Stage Indices.
func checkStageIndices() error {

	var err error

	// Stage Indices.
	if (Stage_Index_Max - Stage_Index_Min + 1) != StagesCount {
		err = errors.New(ErrMsgSelfCheckStageIndices)
		return err
	}

	// Stage Status Indices.
	if (Status_Index_Max - Status_Index_Min + 1) != StatusesCount {
		err = errors.New(ErrMsgSelfCheckStatusIndices)
		return err
	}

	return nil
}
