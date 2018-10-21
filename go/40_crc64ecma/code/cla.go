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

// cla.go.

// CRC 64 ECMA :: Command Line Arguments.

package main

import "flag"

// Common Parameters.
const CLAFullNamePrefix = "-"

// Argument Names.
const CLAName_ResultUpperCase = "uc"
const CLAFullName_ResultUpperCase = CLAFullNamePrefix + CLAName_ResultUpperCase

// Lengths of Argument full Names.
const CLAFullNameLen_ResultUpperCase = 3 // '-uc'.

// Argument Default Values.
const CLADefaultValue_ResultUpperCase = false

// Argument Hint Texts.
const CLAHint_ResultUpperCase = "Use Upper Case Letters in Result"

// Application Configuration Parameters taken from Command Line Arguments.
var upperCaseLettersInResultAreUsed bool

// Initializes Command Line Arguments.
func claInit() error {

	var pUpperCaseLettersInResultAreUsed *bool

	// Upper Case Letters in Result.
	pUpperCaseLettersInResultAreUsed = flag.Bool(
		CLAName_ResultUpperCase,
		CLADefaultValue_ResultUpperCase,
		CLAHint_ResultUpperCase,
	)

	// Read Flags.
	flag.Parse()
	upperCaseLettersInResultAreUsed = *pUpperCaseLettersInResultAreUsed

	return nil
}

// Checks whether the Key is a Configuration Parameter or not.
func claIsAConfigurationParameter(key string) bool {

	// Check Length.
	if len(key) != CLAFullNameLen_ResultUpperCase {
		return false
	}

	// Check Content.
	if key != CLAFullName_ResultUpperCase {
		return false
	}

	return true
}
