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

// main.go.

//	CRC 64 ECMA :: Main File.

//	CRC 64 ECMA.

//	64-Bit Cyclic Redundancy Check using the
//	Polynomial Table provided by the
//	European Computer Manufacturers Association (ECMA).

//	Notes:
//	This Program can read only small Files which can be placed in the RAM.

//	Version:	0.2.

package main

import (
	"fmt"
	"log"
	"os"
)

const MsgUsage = "Usage: crc64ecma <file_1> <file_2> ..."

// Program's Entry Point.
func main() {

	var cla string
	var claCount int
	var err error
	var filesCount int
	var i int
	var sum string

	// Check Command Line Arguments Count.
	claCount = len(os.Args)
	if claCount <= 1 {

		// No Arguments are set.
		fmt.Println(MsgUsage)
		return
	}

	// Initializations.
	err = initialize()
	if err != nil {

		// Initialization Failure.
		log.Println(ErrInitialization)
		return
	}

	// Get File Paths from Command Line Arguments and process each File.
	filesCount = claCount - 1
	for i = 1; i <= filesCount; i++ {

		// Get Argument.
		cla = os.Args[i]

		// Check whether it is a File Path or a Configuration Key.
		if claIsAConfigurationParameter(cla) {

			// Configuration Key is ignored.
			continue
		}

		// Calculate Check Sum.
		sum, err = calculateSumStringHex(cla)
		if err != nil {

			// Error Report.
			log.Println(ErrPrefixCalculation, cla, err)
			continue
		}

		// Successful Result Report.
		fmt.Println(sum, cla)
		continue
	}

	return
}
