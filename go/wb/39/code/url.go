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

// url.go.

// URL Processor.

// Date: 2018-10-19.

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const ErrPrefixHTTP = "HTTP Error. "

const SearchPattern = "Go"

const MsgPrefixProcessingUrl = "Processing URL: "

// Processes an URL.
// On Success, saves the Results.
func urlProcess(
	url string,
	resultsChannel chan ResultEntry,
	verboseMode bool,
) {

	var err error
	var msg string
	var patternsCount int
	var response *http.Response
	var responseBody []byte
	var responseBodyStr string
	var result ResultEntry

	if verboseMode {
		// Log.
		msg = MsgPrefixProcessingUrl + url
		log.Println(msg)
	}

	// Get a Response.
	response, err = http.Get(url)
	if err != nil {
		if verboseMode {
			// Log.
			msg = ErrPrefixHTTP + err.Error()
			log.Println(msg)
		}
		return
	}

	// Get Response Body.
	defer response.Body.Close()
	responseBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		// Log.
		msg = ErrPrefixHTTP + err.Error()
		log.Println(msg)
	}
	responseBodyStr = string(responseBody)

	// Count Patterns in Response.
	patternsCount = strings.Count(responseBodyStr, SearchPattern)

	// Save Results.
	result.URL = url
	result.MatchesCount = patternsCount
	resultsChannel <- result
}
