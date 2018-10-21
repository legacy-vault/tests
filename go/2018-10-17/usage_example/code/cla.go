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

// cla.go.

// CSV File Parsing Example :: Command Line Arguments.

// Command Line Arguments Processing.

// Author: McArcher.
// Date: 2018-10-17.

package main

import "flag"

// Names of Command Line Arguments.
const CLAName_FileInput = "in"
const CLAName_DBAddress = "db_address"
const CLAName_DBAuthDataBase = "db_auth_base"
const CLAName_DBDataBase = "db_data_base"
const CLAName_DBPassword = "db_pwd"
const CLAName_DBUsername = "db_user"

// Hint Texts of Command Line Arguments.
const CLAHintText_FileInput = "Path to the Input File"
const CLAHintText_DBAddress = "Address of the Output Database"
const CLAHintText_DBAuthDataBase = "Authentication Base of the Output Database"
const CLAHintText_DBDataBase = "Base of the Output Database"
const CLAHintText_DBPassword = "Password of the Output Database"
const CLAHintText_DBUsername = "Username of the Output Database"

// Default Values of Command Line Arguments.
const CLADefaultValue_FileInput = CLAValue_Empty
const CLADefaultValue_DBAddress = MongoDBHostDefaultStr +
	MongoDBAddressDelimiterHostPort +
	MongoDBPortDefaultStr
const CLADefaultValue_DBAuthDataBase = CLAValue_Empty
const CLADefaultValue_DBDataBase = CLAValue_Empty
const CLADefaultValue_DBPassword = CLAValue_Empty
const CLADefaultValue_DBUsername = CLAValue_Empty

// Particular Values of various Configuration Parameters.
const CLAValue_Empty = ""
const MongoDBHostDefaultStr = "localhost"
const MongoDBPortDefaultStr = "27017"
const MongoDBAddressDelimiterHostPort = ":"

// Receives input Data from Command Line Arguments.
func claDataReceive(
	inFilePath *string,        // [W]
	outDBAddress *string,      // [W]
	outDBAuthDataBase *string, // [W]
	outDBDataBase *string,     // [W]
	outDBPassword *string,     // [W]
	outDBUsername *string,     // [W]
) error {

	var pInFilePath *string
	var pOutDBAddress *string
	var pOutDBAuthDataBase *string
	var pOutDBDataBase *string
	var pOutDBPassword *string
	var pOutDBUsername *string

	// Configure Flags...

	// 1. Input File Path.
	pInFilePath = flag.String(
		CLAName_FileInput,
		CLADefaultValue_FileInput,
		CLAHintText_FileInput,
	)

	// 2. Output Database Address.
	pOutDBAddress = flag.String(
		CLAName_DBAddress,
		CLADefaultValue_DBAddress,
		CLAHintText_DBAddress,
	)

	// 3. Output Database Authentication Base.
	pOutDBAuthDataBase = flag.String(
		CLAName_DBAuthDataBase,
		CLADefaultValue_DBAuthDataBase,
		CLAHintText_DBAuthDataBase,
	)

	// 4. Output Database Data Base.
	pOutDBDataBase = flag.String(
		CLAName_DBDataBase,
		CLADefaultValue_DBDataBase,
		CLAHintText_DBDataBase,
	)

	// 5. Output Database Password.
	pOutDBPassword = flag.String(
		CLAName_DBPassword,
		CLADefaultValue_DBPassword,
		CLAHintText_DBPassword,
	)

	// 6. Output Database Username.
	pOutDBUsername = flag.String(
		CLAName_DBUsername,
		CLADefaultValue_DBUsername,
		CLAHintText_DBUsername,
	)

	// Parse Flags.
	flag.Parse()

	// Get Values.
	*inFilePath = *pInFilePath
	*outDBAddress = *pOutDBAddress
	*outDBAuthDataBase = *pOutDBAuthDataBase
	*outDBDataBase = *pOutDBDataBase
	*outDBPassword = *pOutDBPassword
	*outDBUsername = *pOutDBUsername

	return nil
}
