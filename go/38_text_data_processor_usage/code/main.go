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

// CSV File Parsing Example :: Main File.

// A simple Example of CSV File parsing Library Usage.

// Author: McArcher.
// Date: 2018-10-17.

package main

import (
	"fmt"
)

const ErrHint = "Error:"
const ErrImport = "Import has failed."
const ErrInitialization = "Initialization has failed."
const ErrInputFilePath = "Input File Path Error"
const ErrOutputDatabaseAddress = "Output Database Address Error"
const ErrOutputDatabaseBase = "Output Database Base Error"
const ErrOutputDatabaseAuthBase = "Output Database Authentication Base Error"
const ErrOutputDatabaseUsername = "Output Database Username Error"

var inputFilePath string
var outputDBAddress string
var outputDBAuthDataBase string
var outputDBDataBase string
var outputDBPassword string
var outputDBUsername string
var outputDBAuthIsRequired bool

// Program's Entry Point.
func main() {

	var err error

	// Initializations.
	err = initialize()
	if err != nil {
		fmt.Println(ErrInitialization)
		fmt.Println(ErrHint, err.Error())
		return
	}

	// Main Work.
	err = work()
	if err != nil {
		fmt.Println(ErrImport)
		fmt.Println(ErrHint, err.Error())
		return
	}
}
