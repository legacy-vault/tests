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

// insert.go.

// Compact Double Link List :: List Object 'Insert' Functions.

package cdllist

// Inserts an Item as a new Head into the List.
func (list *List) InsertHead(item *ListItem) {

	var oldHead *ListItem

	if list.size == 0 {

		// Insert an Item to the empty List.
		item.previousItem = nil
		item.nextItem = nil
		list.head = item
		list.tail = item
		list.size++

		return
	}

	// Insert an Item to the non-empty List.
	oldHead = list.head
	item.nextItem = oldHead
	item.previousItem = nil
	oldHead.previousItem = item
	list.head = item
	list.size++

	return
}

// Inserts an Item after the Reference Item.
// Returns 'true' on Success.
func (list *List) InsertNextItem(
	referenceItem *ListItem,
	newNextItem *ListItem,
) bool {

	var ok bool
	var oldNextItem *ListItem
	var tail *ListItem

	// Check input Parameters.
	if referenceItem == nil {
		return false
	}
	if newNextItem == nil {
		return false
	}

	// Is a Reference Item a Tail?
	tail = list.tail
	if referenceItem == tail {
		list.InsertTail(newNextItem)
		return true
	}

	// Is a Reference Item an Item of our List?
	ok = list.HasAnItem(referenceItem)
	if !ok {
		return false
	}

	// Insert an Item normally.
	oldNextItem = referenceItem.nextItem
	newNextItem.nextItem = oldNextItem
	newNextItem.previousItem = referenceItem
	referenceItem.nextItem = newNextItem
	oldNextItem.previousItem = newNextItem
	list.size++

	return true
}

// Unsafely inserts an Item after the Reference Item.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) InsertNextItemUnsafe(
	referenceItem *ListItem,
	newNextItem *ListItem,
) {

	var oldNextItem *ListItem
	var tail *ListItem

	// Is a Reference Item a Tail?
	tail = list.tail
	if referenceItem == tail {
		list.InsertTail(newNextItem)
		return
	}

	// Insert an Item normally.
	oldNextItem = referenceItem.nextItem
	newNextItem.nextItem = oldNextItem
	newNextItem.previousItem = referenceItem
	referenceItem.nextItem = newNextItem
	oldNextItem.previousItem = newNextItem
	list.size++

	return
}

// Inserts an Item before the Reference Item.
// Returns 'true' on Success.
func (list *List) InsertPreviousItem(
	newPreviousItem *ListItem,
	referenceItem *ListItem,
) bool {

	var head *ListItem
	var ok bool
	var oldPreviousItem *ListItem

	// Check input Parameters.
	if referenceItem == nil {
		return false
	}
	if newPreviousItem == nil {
		return false
	}

	// Is a Reference Item a Head?
	head = list.head
	if referenceItem == head {
		list.InsertHead(newPreviousItem)
		return true
	}

	// Is a Reference Item an Item of our List?
	ok = list.HasAnItem(referenceItem)
	if !ok {
		return false
	}

	// Insert an Item normally.
	oldPreviousItem = referenceItem.previousItem
	newPreviousItem.nextItem = referenceItem
	newPreviousItem.previousItem = oldPreviousItem
	referenceItem.previousItem = newPreviousItem
	oldPreviousItem.nextItem = newPreviousItem
	list.size++

	return true
}

// Unsafely inserts an Item before the Reference Item.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) InsertPreviousItemUnsafe(
	newPreviousItem *ListItem,
	referenceItem *ListItem,
) {

	var head *ListItem
	var oldPreviousItem *ListItem

	// Is a Reference Item a Head?
	head = list.head
	if referenceItem == head {
		list.InsertHead(newPreviousItem)
		return
	}

	// Insert an Item normally.
	oldPreviousItem = referenceItem.previousItem
	newPreviousItem.nextItem = referenceItem
	newPreviousItem.previousItem = oldPreviousItem
	referenceItem.previousItem = newPreviousItem
	oldPreviousItem.nextItem = newPreviousItem
	list.size++

	return
}

// Inserts an Item as a new Tail into the List.
func (list *List) InsertTail(item *ListItem) {

	var oldTail *ListItem

	if list.size == 0 {

		// Insert an Item to the empty List.
		item.previousItem = nil
		item.nextItem = nil
		list.head = item
		list.tail = item
		list.size++

		return
	}

	// Insert an Item to the non-empty List.
	oldTail = list.tail
	item.previousItem = oldTail
	item.nextItem = nil
	oldTail.nextItem = item
	list.tail = item
	list.size++

	return
}
