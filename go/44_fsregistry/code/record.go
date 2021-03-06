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
// Creation Date:	2018-10-21.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// record.go.

// Fixed Size Registry :: Record Object.

package fsregistry

type Record struct {
	// Usable Record Data.
	Data interface{}

	// Time of Creation of this Record. A UNIX Timestamp.
	toc int64
}

// Gets the Creation Time of a Record.
func (record *Record) GetTimeOfCreation() int64 {

	var toc int64

	toc = record.toc

	return toc
}
