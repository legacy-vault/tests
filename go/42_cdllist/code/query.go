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

// query.go.

// Compact Double Link List :: List Object 'Query' Functions.

package cdllist

// Checks whether the specified Item belongs to the List.
// Returns 'true' if an Item is a Part of the List.
// While this Package provides a compact List Data Model,
// we do not store the Pointer to the Owner-List in each List Item.
// This greatly reduces Memory Load for Lists with simple Data Model and
// big Size, but, as a Drawback, we have to do thorough Calculations each
// Time we want to delete an Item from the List.
func (list *List) HasAnItem(item *ListItem) bool {

	var cursor *ListItem
	var i uint64
	var size uint64
	var tail *ListItem

	// Check input Parameter.
	if item == nil {
		return false
	}

	// Prepare Data.
	tail = list.tail
	size = list.size

	// Empty List?
	if size == 0 {
		return false
	}

	// Search from Tail to Head.
	cursor = tail
	i = 1
	for {
		// Found?
		if cursor == item {
			return true
		}
		cursor = cursor.previousItem
		// End of the Line and still not found?
		if cursor == nil {
			return false
		}
		i++
		// Defence against Self-Loop Anomaly.
		if i > size {
			break
		}
	}
	// This Code is reachable only in Case of Self-Loop Anomaly.
	return false
}

// Checks whether the List is not empty.
func (list *List) HasItems() bool {

	if list.size != 0 {
		return true
	}

	return false
}

// Checks whether the List is empty.
func (list *List) IsEmpty() bool {

	if list.size == 0 {
		return true
	}

	return false
}

// Checks the List's Integrity.
// This is a Self-Check Function intended to find Anomalies.
// This Function is not intended to be used in an ordinary Case.
// Returns 'true' if the List is in a good Shape.
func (list *List) IsIntegral() bool {

	var cursor *ListItem
	var cursorNextItem *ListItem
	var cursorPreviousItem *ListItem
	var head *ListItem
	var i uint64
	var size uint64
	var sizeAnomaly bool
	var tail *ListItem

	// Prepare Data.
	head = list.head
	tail = list.tail
	size = list.size

	// Empty List?
	if size == 0 {
		if head != nil {
			return false
		}
		if tail != nil {
			return false
		}
		return true
	}

	// Single-Item List?
	if size == 1 {
		if head == nil {
			return false
		}
		if tail == nil {
			return false
		}
		if head != tail {
			return false
		}
		if head.previousItem != nil {
			return false
		}
		if tail.nextItem != nil {
			return false
		}
		return true
	}

	// List has two or more Items.

	// Check Head Corner.
	if head.previousItem != nil {
		return false
	}

	// Try to inspect all Items from Head to Tail.
	// This checks Connectivity by the 'next' Pointer.
	cursor = head
	cursorNextItem = cursor.nextItem
	i = 1
	for cursorNextItem != nil {
		cursor = cursorNextItem
		cursorNextItem = cursor.nextItem
		i++
		// Defence against Self-Loop Anomaly.
		if i > size {
			sizeAnomaly = true
			break
		}
	}
	if sizeAnomaly {
		// Size Anomaly can happen if we either have a Self-Loop in the Chain
		// or the Corner Item for some Reason is not the End of the Chain.
		return false
	}
	if i != size {
		return false
	}
	// We have stopped the Search at the first Break in the Chain.
	// Are we really there where we should be?
	if cursor != tail {
		// We have found a broken Connection.
		return false
	}

	// Check Tail Corner.
	if tail.nextItem != nil {
		return false
	}

	// Now, try to inspect all Items in a reversed Order.
	// This checks Connectivity by the 'previous' Pointer.
	cursor = tail
	cursorPreviousItem = cursor.previousItem
	i = 1
	for cursorPreviousItem != nil {
		cursor = cursorPreviousItem
		cursorPreviousItem = cursor.previousItem
		i++
		// Defence against Self-Loop Anomaly.
		if i > size {
			sizeAnomaly = true
			break
		}
	}
	if sizeAnomaly {
		// Size Anomaly can happen if we either have a Self-Loop in the Chain
		// or the Corner Item for some Reason is not the End of the Chain.
		return false
	}
	if i != size {
		return false
	}
	// We have stopped the Search at the first Break in the Chain.
	// Are we really there where we should be?
	if cursor != head {
		// We have found a broken Connection.
		return false
	}

	// All Clear.
	return true
}
