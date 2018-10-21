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
// Creation Date:	2018-10-19.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// calculator.go

// CRC 64 ECMA :: Calculator.

package main

import (
	"hash/crc64"
	"io/ioutil"
	"strconv"
	"strings"
)

const ErrPrefixCalculation = "Checksum Calculation Error. File:"

const ResultOnError = ""

var crcTable *crc64.Table

// Calculates the ECMA CRC-64 Check Sum of the File.
// File is specified by its Path.
// Returns the Sum as a Hexadecimal String.
func calculateSumStringHex(filePath string) (string, error) {

	var crcSum uint64
	var err error
	var fileData []byte
	var result string

	// Read the entire File into RAM.
	fileData, err = ioutil.ReadFile(filePath)
	if err != nil {
		return ResultOnError, err
	}

	// Calculate Check Sum.
	crcSum = crc64.Checksum(fileData, crcTable)

	// Uint64 -> String.
	result = strconv.FormatUint(crcSum, 16)

	// Upper Case Letters?
	if upperCaseLettersInResultAreUsed {
		result = strings.ToUpper(result)
	}

	return result, nil
}
