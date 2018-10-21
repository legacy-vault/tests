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

// get.go.

// Fixed Size Registry :: 'Get' Functions.

package fsregistry

import (
	"github.com/legacy-vault/library/go/compact_double_link_list"
	"log"
)

// Gets the actual Size of the Registry,
// the Number of Records actually stored.
func (registry *Registry) GetSize() uint64 {

	var size uint64

	size = registry.records.GetSize()

	return size
}

// Gets the last Record from the Registry (if it is possible)
// and the Registry Size.
func (registry *Registry) GetLastRecord() (*Record, uint64) {

	var listItem *cdllist.ListItem
	var ok bool
	var record *Record
	var size uint64

	// Zero Size?
	size = registry.records.GetSize()
	if size == 0 {
		return nil, 0
	}

	// Get a Record normally.
	listItem = registry.records.GetHead()
	record, ok = listItem.Data.(*Record)
	if !ok {
		// Internal Error, Anomaly.
		log.Println(ErrAnomalyTypeAssertion)
	}

	return record, size
}

// Gets all stored Records from the Registry and their Quantity.
func (registry *Registry) GetStoredRecords() ([]*Record, uint64) {

	var i uint64
	var ok bool
	var recordRaw interface{}
	var records []*Record
	var recordsRaw []interface{}
	var size uint64

	// Prepare Data.
	size = registry.records.GetSize()
	records = make([]*Record, size)

	// Zero Size?
	if size == 0 {
		return records, 0
	}

	// Get all Records normally.
	recordsRaw = registry.records.EnlistAllItemValues()
	for i = 0; i < size; i++ {
		recordRaw = recordsRaw[i]
		records[i], ok = recordRaw.(*Record)
		if !ok {
			// Internal Error, Anomaly.
			log.Println(ErrAnomalyTypeAssertion)
		}
	}

	return records, size
}
