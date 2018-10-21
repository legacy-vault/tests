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

// move.go.

// Compact Double Link List :: List Object 'Move' Functions.

package cdllist

// Moves an Item to a Position after the Reference Item.
// Returns 'true' on Success.
func (list *List) MoveItemToAfterReference(
	referenceItem *ListItem,
	movedItem *ListItem,
) bool {

	var ok bool

	ok = list.RemoveItem(movedItem)
	if !ok {
		return false
	}

	ok = list.InsertNextItem(referenceItem, movedItem)
	if !ok {
		return false
	}
	return true
}

// Unsafely moves an Item to a Position after the Reference Item.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) MoveItemToAfterReferenceUnsafe(
	referenceItem *ListItem,
	movedItem *ListItem,
) {

	list.RemoveItemUnsafe(movedItem)
	list.InsertNextItemUnsafe(referenceItem, movedItem)
}

// Moves an Item to a Position before the Reference Item.
// Returns 'true' on Success.
func (list *List) MoveItemToBeforeReference(
	movedItem *ListItem,
	referenceItem *ListItem,
) bool {

	var ok bool

	ok = list.RemoveItem(movedItem)
	if !ok {
		return false
	}

	ok = list.InsertPreviousItem(movedItem, referenceItem)
	if !ok {
		return false
	}
	return true
}

// Unsafely moves an Item to a Position before the Reference Item.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) MoveItemToBeforeReferenceUnsafe(
	movedItem *ListItem,
	referenceItem *ListItem,
) {

	list.RemoveItemUnsafe(movedItem)
	list.InsertPreviousItemUnsafe(movedItem, referenceItem)
}

// Moves an Item to the Head Position.
// Returns 'true' on Success.
func (list *List) MoveItemToHeadPosition(item *ListItem) bool {

	var ok bool

	ok = list.RemoveItem(item)
	if !ok {
		return false
	}

	list.InsertHead(item)
	return true
}

// Unsafely moves an Item to the Head Position.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) MoveItemToHeadPositionUnsafe(item *ListItem) {

	list.RemoveItemUnsafe(item)
	list.InsertHead(item)
}

// Moves an Item to the Tail Position.
// Returns 'true' on Success.
func (list *List) MoveItemToTailPosition(item *ListItem) bool {

	var ok bool

	ok = list.RemoveItem(item)
	if !ok {
		return false
	}

	list.InsertTail(item)
	return true
}

// Unsafely moves an Item to the Tail Position.
// This Function is for those old-school Punks who like the C Style.
// Works fast but all the Checks are skipped.
func (list *List) MoveItemToTailPositionUnsafe(item *ListItem) {

	list.RemoveItemUnsafe(item)
	list.InsertTail(item)
}
