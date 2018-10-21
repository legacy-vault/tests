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

// item.go.

// Compact Double Link List :: List Item Object Functions.

package cdllist

type ListItem struct {
	// Public Parameters.
	Data interface{}

	// Internal Parameters.
	nextItem     *ListItem
	previousItem *ListItem
}

// Returns the next List Item.
func (item *ListItem) GetNextItem() *ListItem {

	var nextItem *ListItem

	nextItem = item.nextItem

	return nextItem
}

// Returns the previous List Item.
func (item *ListItem) GetPreviousItem() *ListItem {

	var previousItem *ListItem

	previousItem = item.previousItem

	return previousItem
}

// Checks whether an Item has next Item.
func (item *ListItem) HasNextItem() bool {

	var nextItem *ListItem

	nextItem = item.nextItem
	if nextItem != nil {
		return true
	}

	return false
}

// Checks whether an Item has no next Item.
func (item *ListItem) HasNoNextItem() bool {

	var nextItem *ListItem

	nextItem = item.nextItem
	if nextItem == nil {
		return true
	}

	return false
}

// Checks whether an Item has no previous Item.
func (item *ListItem) HasNoPreviousItem() bool {

	var previousItem *ListItem

	previousItem = item.previousItem
	if previousItem == nil {
		return true
	}

	return false
}

// Checks whether an Item has previous Item.
func (item *ListItem) HasPreviousItem() bool {

	var previousItem *ListItem

	previousItem = item.previousItem
	if previousItem != nil {
		return true
	}

	return false
}
