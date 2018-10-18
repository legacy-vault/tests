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

// check.go.

// CSV File Parsing Example :: Data Check.

// Input Data Check.

// Author: McArcher.
// Date: 2018-10-17.

package main

import (
	"errors"
)

// Checks Parameters.
// Sets the 'outDBAuthIsRequired' Parameter to 'true'
// if input Parameters imply the User Authentication.
func inputParametersCheck(
	inFilePath string,         // [R]
	outDBAddress string,       // [R]
	outDBDataBase string,      // [R]
	outDBAuthIsRequired *bool, // [W]
	outDBAuthDataBase string,  // [R]
	outDBUsername string,      // [R]
) error {

	var err error
	var inFilePathLen int
	var outDBAddressLen int
	var outDBAuthDataBaseLen int
	var outDBDataBaseLen int
	var outDBUsernameLen int

	// Input File Path Check.
	inFilePathLen = len(inFilePath)
	if (inFilePathLen == 0) {
		err = errors.New(ErrInputFilePath)
		return err
	}

	// Output Database Address Check.
	outDBAddressLen = len(outDBAddress)
	if (outDBAddressLen == 0) {
		err = errors.New(ErrOutputDatabaseAddress)
		return err
	}

	// Output Database Base Check.
	outDBDataBaseLen = len(outDBDataBase)
	if (outDBDataBaseLen == 0) {
		err = errors.New(ErrOutputDatabaseBase)
		return err
	}

	// If any of the below listed Parameters are set (not empty),
	// then Database requires Authentication:
	//	-	DB Authentications DataBase,
	//	-	DB Username.
	outDBAuthDataBaseLen = len(outDBAuthDataBase)
	outDBUsernameLen = len(outDBUsername)
	if (outDBAuthDataBaseLen > 0) || (outDBUsernameLen > 0) {
		*outDBAuthIsRequired = true
	}
	if (*outDBAuthIsRequired == true) {
		if (outDBAuthDataBaseLen == 0) {
			err = errors.New(ErrOutputDatabaseAuthBase)
			return err
		}
		if (outDBUsernameLen == 0) {
			err = errors.New(ErrOutputDatabaseUsername)
			return err
		}
	}

	return nil
}
