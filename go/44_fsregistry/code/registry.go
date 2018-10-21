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

// registry.go.

// Fixed Size Registry :: Registry Object Constructor.

// Stores Information about chronologically last N Records,
// where 'N' is a fixed Number of Records.
// Works as a simple 'First In - First Out' (FIFO) Queue of a fixed Size.

package fsregistry

import (
	"github.com/legacy-vault/library/go/compact_double_link_list"
)

type Registry struct {
	records  cdllist.List
	capacity uint64
}

// Creates a new Registry and returns a Pointer directed at it.
func New(capacity uint64) *Registry {

	var registry *Registry

	registry = new(Registry)
	registry.initialize(capacity)

	return registry
}

// Initializes the Registry Object.
func (registry *Registry) initialize(capacity uint64) {

	var list *cdllist.List

	registry.capacity = capacity
	list = cdllist.CreateNewList()
	registry.records = *list
}
