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

// init.go.

// Main Initializations.

// Date: 2018-10-19.

package main

// Initializes the Application.
func initialize(app *Application) error {

	var err error

	// Read Command Line Arguments.
	err = claInit(app)
	if err != nil {
		return err
	}

	// Enable Reading from Standard Input Stream.
	app.ReadingFromStdinIsEnabled = true

	// Prepare Control Channels.
	app.WorkersQuitChannel = make(chan bool, WorkersQuitChannelSize)
	app.ResultsCollectorQuitChan = make(chan bool, ResultsCollectorQuitChan)

	// Initialize Data Channels.
	app.RawDataChannel = make(chan string, RawDataChannelSize)
	app.ResultsChannel = make(chan ResultEntry, ResultsChannelSize)

	// Initialize Results Storage.
	app.Results = make([]ResultEntry, 0)

	// Initialize the Operation System Signal Handler.
	err = ossInit(app)
	if err != nil {
		return err
	}

	return nil
}
