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
// Creation Date:	2018-10-17.
// Web Site Address is an Address in the global Computer Internet Network.
//
//============================================================================//

// model.go.

// CSV File Parsing Example :: Object Model.

// Model of an Object used as Reference for Data Parsing.

// Author: McArcher.
// Date: 2018-10-17.

package main

// Planet Object Model.
type Planet struct {
	// Unique Identifier (ID) of a Planet's Parent Object (a single Star or
	// another Object) in a Space Catalog.
	ParentUID string

	// Distance from the Parent Object measured in Kilometers.
	ParentDistanceKm uint64

	// Unique Identifier (ID) of a Planet in a Space Catalog.
	UID string

	// Visible (by Human Eye) Diameter in Kilometers.
	DiameterKm uint64

	// Life Stage Relative Measurement Parameter.
	LSRMP float64

	// Is Planet a Member of Inter-Galactic Union 'IGC-I'?
	IsAMemberOfIGCUno bool
}
