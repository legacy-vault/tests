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

// result.go.

// Results Storing.

// Date: 2018-10-19.

package main

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

// Messages.
const MsgResultsChannelClosed = "Results Channel Read Failure. " +
	"Channel is closed."
const MsgResultsCollectorStop = "Results Collector has been stopped."

// Message Formats.
const MsgFormatResultsCollectorStatistics = "Results Collector " +
	"has received %v Results.\r\n"

const SleepTickMsResultsCollector = 250

type ResultEntry struct {
	URL          string
	MatchesCount int
}

// Collects the Results.
func resultsCollector(
	resultsChannel chan ResultEntry,
	results *[]ResultEntry,
	busyCollector *sync.WaitGroup,
	collectorQuitChannel chan bool,
	verboseMode bool,
) {

	var loop bool
	var received bool
	var receivedResultsCount uint64
	var result ResultEntry
	var sleepTick time.Duration
	var stopSignal bool

	defer busyCollector.Done()

	// Preparations.
	sleepTick = time.Millisecond * SleepTickMsResultsCollector

	// Receive raw Data.
	loop = true
	for loop {

		select {

		case result, received = <-resultsChannel:
			if received == false {

				// Channel is not available.

				// Log.
				log.Println(MsgResultsChannelClosed)
				loop = false

			} else {

				// New incoming Result.
				receivedResultsCount++

				// Save it.
				*results = append(*results, result)
			}

		default:
			// No Results.
			// Such Situation can happen in two Cases:
			//	1.	We are using manual (Keyboard) Input and the User has not
			// 		yet put anything into the Stream.
			//	2.	The Stream Reader has been stopped and the Queue is empty.

			// Try to get a Stop Signal.
			select {

			case stopSignal = <-collectorQuitChannel:

				if stopSignal == ResultsCollectorQuitChanSignalQuit {
					// The Queue is empty and we must quit.
					loop = false
				}

			default:

				// No Tasks are available & no Stop Signal.
				// Sleep for a while...
				time.Sleep(sleepTick)
			}
		}
	}

	// Verbose Log.
	if verboseMode {
		log.Printf(MsgFormatResultsCollectorStatistics, receivedResultsCount)
		log.Println(MsgResultsCollectorStop)
	}
}

// Shows the Results.
func resultsShow(results []ResultEntry) {

	var i int
	var overflowHappened bool
	var result ResultEntry
	var quantityTotal uint64
	var quantityTotalDelta uint64

	fmt.Println("------------[RESULTS]------------------------------------")
	fmt.Println("|      #. | Match Qty. |    URL                         |")
	fmt.Println("---------------------------------------------------------")
	i = 1
	for _, result = range results {
		fmt.Printf(
			"| %6d. | %10d | %s \r\n",
			i,
			result.MatchesCount,
			result.URL,
		)
		i++

		// Calculate Total Quantity.
		quantityTotalDelta = uint64(result.MatchesCount)
		if quantityTotalDelta > (math.MaxUint64 - quantityTotal) {
			overflowHappened = true
		}
		quantityTotal += quantityTotalDelta
	}
	fmt.Println("---------------------------------------------------------")
	if !overflowHappened {
		// No Overflow.
		// Notes: Maximum unsigned 64-bit Integer Number has 20 Decimal Digits.
		fmt.Printf(
			"| Total Matches Quantity: %20d. \r\n",
			quantityTotal,
		)
	} else {
		// Overflow.
		fmt.Printf(
			"| Total Matches Quantity: OVERFLOW. \r\n",
		)
	}
	fmt.Println("---------------------------------------------------------")
}
