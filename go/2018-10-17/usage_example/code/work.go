//============================================================================//
//
// Copyright © 2018 by McArcher.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
//============================================================================//
//
// Web Site:		'https://github.com/legacy-vault'.
// Author:			McArcher.
// Creation Date:	2018-10-17.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// work.go.

// Main Work Function.

// Author: McArcher.
// Date: 2018-10-17.

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	tdp "github.com/legacy-vault/library/go/text_data_processor"
)

const CSVFieldSeparator = ';'

// Processes Data from input File.
func work() error {

	var cell string
	var columnsCount int
	var columnIdx int
	var columnNames tdp.Names
	var csvReader *csv.Reader
	var err error
	var inputFile *os.File
	var msg string
	var processor tdp.Processor
	var referenceObject Planet
	var row tdp.Cells
	var rowsCount int

	// User Greeting.
	fmt.Printf("Input File: '%s'.\r\n", inputFilePath)

	// Open the input File.
	inputFile, err = os.Open(inputFilePath)
	if err != nil {
		return err
	}

	// Configure the Go's built-in CSV Parser.
	// Fields Count Setting enables Error Return in Case of
	// Fields Count Mismatch.
	csvReader = csv.NewReader(inputFile)
	csvReader.Comma = CSVFieldSeparator
	csvReader.FieldsPerRecord = 0 // => Automatic Mode.
	csvReader.LazyQuotes = false
	csvReader.ReuseRecord = false
	csvReader.TrimLeadingSpace = false

	// Read the first Row to get the List of Column Names.
	rowsCount = -1
	row, err = csvReader.Read()
	if err != nil {
		return err
	}
	rowsCount++
	columnsCount = len(row)

	// Get Column Names from the first Row.
	columnNames = make(tdp.Names, columnsCount)
	for columnIdx, cell = range row {

		// Get Cell with Column's Name.
		cell = strings.TrimSpace(cell)
		if len(cell) == 0 {
			err = errors.New("Columns Name is empty")
			return err
		}

		columnNames[columnIdx] = cell
	}

	// Configure the Processor.
	err = processor.Configure(&referenceObject, columnNames)
	if err != nil {
		return err
	}

	// Check Input CSV File for Errors.
	// Read all the rest Rows.
	for {

		// Get a Row from File.
		row, err = csvReader.Read()
		if err != nil {

			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		// Process the Row (Row → Object).
		err = processor.Process(row, &referenceObject)
		if err != nil {
			return err
		}

		rowsCount++
	}

	// Congratulations.
	msg = "Input File '%v' has been checked. It is now safe to proceed.\r\n"
	fmt.Printf(msg, inputFilePath)

	// Re-Open the File.
	err = inputFile.Close()
	if err != nil {
		return err
	}
	inputFile, err = os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	// Re-Initialize the Reader.
	csvReader = csv.NewReader(inputFile)
	csvReader.Comma = CSVFieldSeparator
	csvReader.FieldsPerRecord = 0 // => Automatic Mode.
	csvReader.LazyQuotes = false
	csvReader.ReuseRecord = false
	csvReader.TrimLeadingSpace = false

	// Skip the first Line of File.
	rowsCount = -1
	_, err = csvReader.Read()
	if err != nil {
		return err
	}
	rowsCount++

	// Prepare Output Database Connection and Structures.
	err = dbPrepare(
		outputDBAddress,
		outputDBDataBase,
		outputDBAuthIsRequired,
		outputDBAuthDataBase,
		outputDBUsername,
		outputDBPassword,
	)
	if err != nil {
		return err
	}

	// Read all the rest Rows and process them.
	for {

		// Get a Row from File.
		row, err = csvReader.Read()
		if err != nil {

			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		// Process the Row (Row → Object).
		err = processor.Process(row, &referenceObject)
		if err != nil {
			return err
		}

		// Insert the Row Data from Object into Output Database.
		err = planetInsertIntoDB(referenceObject)
		if err != nil {
			return err
		}

		rowsCount++
	}

	// Data Table Information Message.
	msg = "Records have been inserted into Base. Have a good Day!"
	fmt.Println(rowsCount, msg)

	return nil
}
