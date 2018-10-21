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

// init.go.

// CSV File Parsing Example :: Program Initialization.

// Program Initialization.

// Author: McArcher.
// Date: 2018-10-17.

package main

// Initializes the Program.
func initialize() error {

	var err error

	// Get Data from CLA.
	err = claDataReceive(
		&inputFilePath,
		&outputDBAddress,
		&outputDBAuthDataBase,
		&outputDBDataBase,
		&outputDBPassword,
		&outputDBUsername,
	)
	if err != nil {
		return err
	}

	// Check input Parameters.
	err = inputParametersCheck(
		inputFilePath,
		outputDBAddress,
		outputDBDataBase,
		&outputDBAuthIsRequired,
		outputDBAuthDataBase,
		outputDBUsername,
	)
	if err != nil {
		return err
	}

	return nil
}
