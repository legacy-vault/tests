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

// remove.go.

// Compact Double Link List :: List Object 'Remove' Functions.

package cdllist

// Deletes all Items from an existing List.
// Returns 'true' on Success.
func ClearList(list *List) bool {

	var i uint64
	var listIsIntegral bool
	var size uint64

	// Before deleting the List, we must ensure that it is not broken.
	// Broken List Deletion would have caused us a lot of Memory Leaks!
	listIsIntegral = list.IsIntegral()
	if !listIsIntegral {
		return false
	}

	// Prepare Data.
	size = list.size

	// Delete all Items from Tail to Head.
	for i = 1; i <= size; i++ {
		list.RemoveTail()
	}

	return true
}

// Removes a Head from the List.
func (list *List) RemoveHead() {

	var newHead *ListItem
	var oldHead *ListItem

	if list.size == 0 {
		return
	}

	if list.size == 1 {

		list.head = nil
		list.tail = nil
		list.size--

		return
	}

	// Delete an Item normally.
	oldHead = list.head
	newHead = oldHead.nextItem
	list.head = newHead
	newHead.previousItem = nil
	oldHead.previousItem = nil
	oldHead.nextItem = nil
	list.size--

	return
}

// Removes an Item from the List.
// Returns 'true' on Success.
func (list *List) RemoveItem(item *ListItem) bool {

	var head *ListItem
	var left *ListItem
	var itemBelongsToList bool
	var right *ListItem
	var tail *ListItem

	// Check input Parameter.
	if item == nil {
		return false
	}

	// Head?
	head = list.head
	if item == head {
		list.RemoveHead()
		return true
	}

	// Tail?
	tail = list.tail
	if item == tail {
		list.RemoveTail()
		return true
	}

	// Check whether the removed Item really belongs to the List.
	// While this Package provides a compact List Data Model,
	// we do not store the Pointer to the Owner-List in each List Item.
	// This greatly reduces Memory Load for Lists with simple Data Model and
	// big Size, but, as a Drawback, we have to do thorough Calculations each
	// Time we want to delete an Item from the List.
	itemBelongsToList = list.HasAnItem(item)
	if !itemBelongsToList {
		return false
	}

	// If an Item is neither Head, nor Tail,
	// it must be linked with Neighbours.
	left = item.previousItem
	if left == nil {
		return false
	}
	right = item.nextItem
	if right == nil {
		return false
	}

	// Delete an Item normally.
	left.nextItem = right
	right.previousItem = left
	item.previousItem = nil
	item.nextItem = nil
	list.size--

	return true
}

// Unsafely removes an Item from the List.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) RemoveItemUnsafe(item *ListItem) {

	var head *ListItem
	var left *ListItem
	var right *ListItem
	var tail *ListItem

	// Head?
	head = list.head
	if item == head {
		list.RemoveHead()
	}

	// Tail?
	tail = list.tail
	if item == tail {
		list.RemoveTail()
	}

	// Prepare Data.
	left = item.previousItem
	right = item.nextItem

	// Delete an Item normally.
	left.nextItem = right
	right.previousItem = left
	item.previousItem = nil
	item.nextItem = nil
	list.size--
}

// Removes a Tail from the List.
func (list *List) RemoveTail() {

	var newTail *ListItem
	var oldTail *ListItem

	if list.size == 0 {
		return
	}

	if list.size == 1 {

		list.head = nil
		list.tail = nil
		list.size--

		return
	}

	// Delete an Item normally.
	oldTail = list.tail
	newTail = oldTail.previousItem
	list.tail = newTail
	newTail.nextItem = nil
	oldTail.previousItem = nil
	oldTail.nextItem = nil
	list.size--

	return
}
