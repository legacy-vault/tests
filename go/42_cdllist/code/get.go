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

// Compact Double Link List :: List Object 'Get' Functions.

package cdllist

// Enlists Values of all Items of the List.
func (list *List) EnlistAllItemValues() []interface{} {

	var i uint64
	var items []*ListItem
	var size uint64
	var values []interface{}

	items = list.EnlistAllItems()
	size = list.size
	values = make([]interface{}, size)
	for i = 0; i < size; i++ {
		values[i] = items[i].Data
	}

	return values
}

// Enlists all Items of the List.
func (list *List) EnlistAllItems() []*ListItem {

	var i uint64
	var item *ListItem
	var items []*ListItem
	var size uint64

	size = list.size
	items = make([]*ListItem, size)
	if size == 0 {
		return items
	}

	// Get the first Item.
	item = list.head
	items[0] = item

	// Get all other Items.
	for i = 1; i < size; i++ {
		item = item.nextItem
		items[i] = item
	}

	return items
}

// Returns List's Head.
func (list *List) GetHead() *ListItem {

	var head *ListItem

	head = list.head

	return head
}

// Returns List's Size.
func (list *List) GetSize() uint64 {

	var size uint64

	size = list.size

	return size
}

// Returns List's Tail.
func (list *List) GetTail() *ListItem {

	var tail *ListItem

	tail = list.tail

	return tail
}
