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

// app.go.

// Main Application Data.

// Date: 2018-10-19.

package main

import "os"

// Workers.
const WorkersCount = 5

// Channel Configuration...

// 1. Size.
const OsSignalsTermChanSize = 1
const AppQuitChanSize = 1
const RawDataChannelSize = 1024 * 1024
const ResultsChannelSize = RawDataChannelSize
const ResultsCollectorQuitChan = 1
const WorkersQuitChannelSize = WorkersCount

// 2. Signals.
const AppQuitChanSignalQuit = true
const WorkersQuitChanSignalQuit = true
const ResultsCollectorQuitChanSignalQuit = true

type Application struct {
	// Control Channels...

	// 1. Application Quit Signals Channel.
	AppQuitChannel chan bool

	// 2. Quit Signals Channel for all Workers.
	WorkersQuitChannel chan bool

	// 3. Quit Signals Channel for Results Collector.
	ResultsCollectorQuitChan chan bool

	// Control Flags...

	// 1. Is Reading from Standard Input Stream enabled?
	ReadingFromStdinIsEnabled bool

	// Data Channels...

	// 1. Channel accepting O.S. Signals for Service Termination.
	OsSignalsTermChan chan os.Signal

	// 2. Channel with raw text Data.
	RawDataChannel chan string

	// 3. Channel with Results.
	ResultsChannel chan ResultEntry

	// Results Storage.
	Results []ResultEntry

	// Application Settings...

	// 1. Verbose Mode Switch.
	VerboseMode bool

	// 2. Keyboard Input Mode Switch.
	KeyboardInputMode bool
}
