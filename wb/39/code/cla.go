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

// Command Line Arguments.

// Date: 2018-10-19.

package main

import "flag"

// Argument Names.
const CLAName_KeyboardInput = "k"
const CLAName_Verbose = "v"

// Argument Default Values.
const CLADefaultValue_KeyboardInput = false
const CLADefaultValue_Verbose = false

// Argument Hint Texts.
const CLAHint_KeyboardInput = "Keyboard Input"
const CLAHint_Verbose = "Verbose Mode"

func claInit(app *Application) error {

	var keyboardInputIsUsed *bool
	var verboseModeIsUsed *bool

	// Keyboard Input Flag.
	keyboardInputIsUsed = flag.Bool(
		CLAName_KeyboardInput,
		CLADefaultValue_KeyboardInput,
		CLAHint_KeyboardInput,
	)

	// Verbose Mode Flag.
	verboseModeIsUsed = flag.Bool(
		CLAName_Verbose,
		CLADefaultValue_Verbose,
		CLAHint_Verbose,
	)

	// Read Flags.
	flag.Parse()
	app.KeyboardInputMode = *keyboardInputIsUsed
	app.VerboseMode = *verboseModeIsUsed

	return nil
}
