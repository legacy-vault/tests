//============================================================================//
//
// Copyright © 2019 by McArcher.
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
// Creation Date:	2019-11-16.
//
//============================================================================//

package processor

// The Task which is used by the Workers.
// Tasks are put into the Queue by the Processor and are taken from the Queue
// by all the active Workers.
type Task struct {
	UrlAddress string
	Result     TaskResult
}

// Results of the Task Procession.
type TaskResult struct {
	NumberOfGoStrings int
}
