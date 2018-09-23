// main.go.

// Actions List.

// Date: 2018-09-22.

package main

import (
	"fmt"
	"test/block_crc/source/bcrc"
)

// Action Indices.
const Action_Unknown = 0
const Action_CreateBlockCRCSum = 1
const Action_CheckCRCSum = 2

// Action Names.
const ActionName_Unknown = ""
const ActionName_CreateBlockCRCSum = "create"
const ActionName_CheckCRCSum = "check"

var actionType int
var actionStage int

// Sets Action Index according to Action String.
func parseAction(actionStr string) {

	switch actionStr {

	case ActionName_CreateBlockCRCSum:
		actionType = Action_CreateBlockCRCSum
		actionStage = Stage_BlockCRCCreation

	case ActionName_CheckCRCSum:
		actionType = Action_CheckCRCSum
		actionStage = Stage_BlockCRCCheck

	default:
		actionType = Action_Unknown
		actionStage = Stage_Finalization
	}
}

// Creates Block CRC Sum.
func bcrcCreate() error {

	var err error

	err = bcrc.BCRCCreate(*claDataFile, claSumFile2, *claBlockSize)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// Checks Block CRC Sum.
func bcrcCheck() error {

	var err error

	err = bcrc.BCRCCheck(*claDataFile, claSumFile2)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
