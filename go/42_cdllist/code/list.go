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

// list.go.

// Compact Double Link List :: List Object 'Create' Functions.

// This Package provides a double Linked List Functionality.

// The List uses a compact Data Model. This means that, as opposed to the
// Golang's built-in 'list' Package, this Package does not store the Pointer
// to the Owner-List in each List Item. This greatly reduces Memory Load for
// Lists that have simple Data Model and big Size, but, as a Drawback, we
// have to do thorough Calculations each Time we want to delete an Item from
// the List (when the deleted Item is neither Head, nor Tail of the List).
// However, it is important to say that Tail and Head Items are deleted in a
// very fast Manner.

// The Package provides extended Functionality comparing with the Golang's
// built-in 'list' Package.

package cdllist

type List struct {
	// Internal Parameters.
	head *ListItem
	size uint64
	tail *ListItem
}

// Creates a new List.
func CreateNewList() *List {

	var list *List

	list = new(List)
	list.initialize()

	return list
}

// Initializes the List.
func (list *List) initialize() {

	list.head = nil
	list.tail = nil
	list.size = 0

	return
}
