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

// functions.go.

// Text Data Processor :: Auxiliary Functions.

// Various auxiliary Functions used in Data Processor.

// Author: McArcher.
// Date: 2018-10-17.

package processor

import (
	"errors"
	"fmt"
	"reflect"
)

// Creates a "routing" Table.
func createColumnIdxToObjectFieldIdxRouting(
	columnNamesFFA Names,
	referenceObjectFieldNamesFFA Names,
) (RoutingTable, error) {

	var columnIdx int
	var columnNameFFA string
	var columnsCount int
	var err error
	var fieldIdx int
	var routingTable RoutingTable

	// Prepare Data.
	columnsCount = len(columnNamesFFA)
	routingTable = make(RoutingTable, columnsCount)

	// Find Field Indices for each simplified Column Name.
	for columnIdx, columnNameFFA = range columnNamesFFA {

		// Find an Index.
		fieldIdx, err = indexOf(columnNameFFA, referenceObjectFieldNamesFFA)
		if err != nil {
			return nil, err
		}

		// Store an Index.
		routingTable[columnIdx] = fieldIdx
	}

	return routingTable, nil
}

// Returns an Index of a Name (String) in the List of Names (Strings).
func indexOf(
	itemRequested string,
	list Names,
) (int, error) {

	var err error
	var i int
	var msg string
	var name string

	// Find the requested Name.
	for i, name = range list {
		if name == itemRequested {
			return i, nil
		}
	}

	// Ooops! No Name has been found.
	msg = fmt.Sprintf(
		ErrFormatIndexOfItem,
		itemRequested,
	)
	err = errors.New(msg)

	return IndexOfIndexOnFailure, err
}

// Lists all first-level Fields of an Object.
func list1stLevelFields(obj EmptyInterface) (Names) {

	var fieldsCount int
	var fieldName string
	var fieldNames Names
	var i int
	var target reflect.Value

	// Prepare Data.
	target = reflect.ValueOf(obj).Elem()
	fieldsCount = target.NumField()
	fieldNames = make(Names, 0)

	// Inspect all Fields of an Object.
	for i = 0; i < fieldsCount; i++ {
		fieldName = target.Type().Field(i).Name
		fieldNames = append(fieldNames, fieldName)
	}

	return fieldNames
}
