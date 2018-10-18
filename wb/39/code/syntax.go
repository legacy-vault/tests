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

// syntax.go.

// Syntax Checker.

// Date: 2018-10-19.

package main

import (
	"errors"
	"net/url"
	"strings"
)

const ResultURLNone = ""

const URLSchemeAllowedHTTP = "http"
const URLSchemeAllowedHTTPS = "https"

const ErrURL = "Bad URL"

// Checks URL Syntax.
// Returns a valid URL (if it exists) and an Error.
func syntaxCheck(rawText string) (string, error) {

	var err error
	var urlParsed *url.URL
	var urlScheme string

	// Trim Junk Characters from raw Text.
	rawText = strings.TrimSpace(rawText)

	// Parse the raw Text into URL.
	urlParsed, err = url.Parse(rawText)
	if err != nil {
		return ResultURLNone, err
	}

	// Check URL Scheme.
	urlScheme = urlParsed.Scheme
	if (urlScheme != URLSchemeAllowedHTTP) &&
		(urlScheme != URLSchemeAllowedHTTPS) {

		// Unfortunately,
		// as it is stated in the Task, I am not allowed to use global
		// Variables, so I can not cache the Error and have to create
		// a new Error every Time it occurs! This is a Waste of CPU Time!
		err = errors.New(ErrURL)
		return ResultURLNone, err
	}

	return rawText, nil
}
