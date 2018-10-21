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

// add.go.

// Fixed Size Registry :: 'Add' Functions.

package fsregistry

import (
	"github.com/legacy-vault/library/go/compact_double_link_list"
	"log"
	"time"
)

// Adds a Record to the Registry.
func (registry *Registry) AddARecord(record *Record) bool {

	var capacity uint64
	var listItem *cdllist.ListItem
	var size uint64

	// Prepare Data.
	size = registry.records.GetSize()
	capacity = registry.capacity
	record.toc = time.Now().Unix()
	listItem = &cdllist.ListItem{Data: record}

	// Insert a Record into the List of Records.
	if size == capacity {

		// Insertion with preliminary Tail Removal.
		registry.records.RemoveTail()
		registry.records.InsertHead(listItem)
		return true

	} else if size < capacity {

		// Insertion without preliminary Tail Removal.
		registry.records.InsertHead(listItem)
		return true

	} else { // size > capacity

		// Internal Error, Anomaly.
		log.Println(ErrAnomalySize)
		return false
	}
}
