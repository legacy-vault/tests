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

// main.go.

// Text Data Processor :: Main File.

// The Processor receives input text Data (from a CSV File's Row)
// and tries to put it into an Object.

// Prior to receiving real Data, the Processor must be taught (using an empty
// reference Object) how to recognize the input Data. During the initial
// teaching Stage, the Processor inspects the reference Object (using the
// built-in 'reflect' Package) and creates a "routing" Table.

// The "routing" Table stores Information about CSV File's Column Names and
// the Names of Object's Fields. To be more precise, the "routing" Table stores
// Indices of the File's Columns and Indices of the Object's Fields.

// Additional Notes.

// This Version of Processor supports following Object Field Types:
//	*	Bool,
//	*	Float64,
//	*	Int,
//	*	Int64,
//	*	String,
//	*	Uint,
//	*	Uint64.

package processor

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const EmptyString = ""

// Error Messages.
const ErrIndex = "Index is out of Bounds"
const ErrConfigurationIsNotFinished = "Processor is not configured"

// Error Messages' Formats.
const ErrFormatFieldTypeUnsupported = "Field '%v' has unsupported Type"
const ErrFormatIndexOfItem = "Item '%v' is not found."

const IndexOfIndexOnFailure = -1

// Column Name Separator Symbols allowed.
const ColumnNameSeparatorSymbolDot = "."
const ColumnNameSeparatorSymbolHyphen = "-"
const ColumnNameSeparatorSymbolSpace = " "
const ColumnNameSeparatorSymbolUnderline = "_"

type Processor struct {
	// The original List of Field Names of a reference Object.
	FieldNames Names

	// The modified Version of a previous Parameter,
	// modified for the fast Access ('FFA' means: For Fast Access).
	FieldNamesFFA Names

	// The original List of Column Names of a text (CSV) File.
	ColumnNames Names

	// The modified Version of a previous Parameter,
	// modified for the fast Access ('FFA' means: For Fast Access).
	ColumnNamesFFA Names

	// The "routing" Table. An Array of Indices.
	// The Key in this Array is of 'int' Type and represents an Index of
	// Column of the text (CSV) File. Index is Zero-based.
	// The Value (Item) in this Array is also of 'int' Type and represents an
	// Index of reference Object's Field. Index is Zero-based.
	CSVColIdxToFieldIdx RoutingTable

	// Array of Symbols used as Separators for Column Names and Field Names.
	// These Separators are removed from Names when these Names are modified
	// for fast Access.
	NameSeparatorSymbols Names

	// This Flag show whether the Processor has already been successfully
	// configured or not.
	IsConfigured bool
}

type Cells []string
type EmptyInterface interface{}
type Name string
type Names []string
type RoutingTable []int

// Configures the Processor.
// This Method must be run after the Processor Creation but before any Data
// Processing.
// Prepares:
//	*	Name Separator Symbols,
//	*	A modified Version of text File Column Names List,
//	*	A modified Version of reference Object Field Names List,
//	*	A "routing" Table.
func (p *Processor) Configure(
	referenceObject EmptyInterface,
	csvColumnNames Names,
) (error) {

	var csvColumnNamesFFA Names
	var err error
	var fieldNames Names
	var fieldNamesFFA Names
	var routingTable RoutingTable
	var separators Names

	// Store Column Names.
	p.ColumnNames = csvColumnNames

	// Prepare Name Separator Symbols.
	separators = []string{
		ColumnNameSeparatorSymbolDot,
		ColumnNameSeparatorSymbolHyphen,
		ColumnNameSeparatorSymbolSpace,
		ColumnNameSeparatorSymbolUnderline,
	}
	p.NameSeparatorSymbols = separators

	// Create a List of modified Column Names for fast Synchronization.
	csvColumnNamesFFA = p.namesModifyFFA(csvColumnNames)
	p.ColumnNamesFFA = csvColumnNamesFFA

	// List all 1-st Level Field Names of an Object.
	fieldNames = list1stLevelFields(referenceObject)
	p.FieldNames = fieldNames

	// Prepare the List of modified Field Names.
	fieldNamesFFA = p.namesModifyFFA(fieldNames)
	p.FieldNamesFFA = fieldNamesFFA

	// Create a "routing" Table.
	routingTable, err = createColumnIdxToObjectFieldIdxRouting(
		csvColumnNamesFFA,
		fieldNamesFFA,
	)
	if err != nil {

		// Configuration Failure.
		p.IsConfigured = false
		return err
	}
	p.CSVColIdxToFieldIdx = routingTable

	// Configuration Success.
	p.IsConfigured = true
	return nil
}

// Returns an Array of modified for fast Access Names.
func (p Processor) namesModifyFFA(names Names) Names {

	var i int
	var name string
	var namesModified []string
	var separator string

	// Prepare Data.
	namesModified = make([]string, len(names))

	// Do the Changes.
	for i, name = range names {

		// Remove Separators from Name.
		for _, separator = range p.NameSeparatorSymbols {
			name = strings.Replace(name, separator, EmptyString, -1)
		}

		// Lower the Letter Case.
		name = strings.ToLower(name)

		// Save Results back into the Array.
		namesModified[i] = name
	}

	return namesModified
}

// Processes the Data.
// The Processor must be configured before doing any Data Processing.
// Reads Cells from a Row and tries to put them into the Target Object.
func (p *Processor) Process(
	dataRow Cells,
	targetPointer EmptyInterface,
) error {

	var cell string
	var cellBool bool
	var cellFloat64 float64
	var cellIdx int
	var cellInt64 int64
	var cellUint64 uint64
	var err error
	var fieldIdx int
	var fieldIdxMin int
	var fieldIdxMax int
	var fieldsCount int
	var fieldName string
	var fieldType reflect.Kind
	var fieldValue reflect.Value
	var msg string
	var target reflect.Value

	// Check Configurations Status.
	if (p.IsConfigured == false) {
		err = errors.New(ErrConfigurationIsNotFinished)
		return err
	}

	// Prepare Data.
	target = reflect.ValueOf(targetPointer).Elem()
	fieldsCount = target.NumField()
	fieldIdxMin = 0
	fieldIdxMax = fieldsCount - 1

	// Inspect all Cells (and Fields).
	for cellIdx, cell = range dataRow {

		// Get Field's Index from "routing" Table using the Column Index.
		fieldIdx = p.CSVColIdxToFieldIdx[cellIdx]

		// Check the Field Index.
		if (fieldIdx < fieldIdxMin) || (fieldIdx > fieldIdxMax) {
			err = errors.New(ErrIndex)
			return err
		}

		// Get Field's Attributes.
		fieldValue = target.Field(fieldIdx)
		fieldType = fieldValue.Kind()
		fieldName = target.Type().Field(fieldIdx).Name

		// Put a Cell's Value into the Object's Field (if it is possible).
		// While there is no Method which satisfies all the Needs,
		// we must use a different Method for each Type of the Object's Field.
		switch fieldType {

		case reflect.String:
			fieldValue.SetString(cell)

		case reflect.Int:
			cellInt64, err = strconv.ParseInt(cell, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(cellInt64)

		case reflect.Int64:
			cellInt64, err = strconv.ParseInt(cell, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetInt(cellInt64)

		case reflect.Float64:
			cellFloat64, err = strconv.ParseFloat(cell, 64)
			if err != nil {
				return err
			}
			fieldValue.SetFloat(cellFloat64)

		case reflect.Uint:
			cellUint64, err = strconv.ParseUint(cell, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetUint(cellUint64)

		case reflect.Uint64:
			cellUint64, err = strconv.ParseUint(cell, 10, 64)
			if err != nil {
				return err
			}
			fieldValue.SetUint(cellUint64)

		case reflect.Bool:
			cellBool, err = strconv.ParseBool(cell)
			if err != nil {
				return err
			}
			fieldValue.SetBool(cellBool)

		default:
			msg = fmt.Sprintf(
				ErrFormatFieldTypeUnsupported,
				fieldName,
			)
			err = errors.New(msg)
			return err
		}
	}

	return nil
}
