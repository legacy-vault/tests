//============================================================================//
//
// Copyright © 2019 by McArcher.
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
// Creation Date:	2019-11-16.
//
//============================================================================//

package main

// Notes.
//
// Read the 'go_tz.txt' File.
//
// To build this Program in old Versions of Go Language (1.11 or 1.12):
// 		GO111MODULE="on" go build
// To build this Program in modern Go Language (Version 1.13):
// 		go build
// Usage Example in Linux OS:
//		echo -e 'microsoft.com\nhttps://yandex.ru\nping.eu\nhttp://ya.ru\nmail.ru\nqwerty.ru' | ./<executable>

import (
	"stdin_url_reader/processor"
)

func main() {

	var streamProcessor processor.Processor

	streamProcessor.Use()
}
