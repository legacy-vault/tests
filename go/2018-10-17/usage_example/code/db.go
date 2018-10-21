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

// db.go.

// CSV File Parsing Example :: Database.

// Database related Stuff.

// Author: McArcher.
// Date: 2018-10-17.

package main

import (
	"fmt"
	"github.com/globalsign/mgo"
)

const MongoDBCollectionName = "test_collection_1"

var dbSession *mgo.Session
var dbDatabase string
var dbCollection *mgo.Collection

// Prepares the Output Database Connection.
func dbPrepare(
	outDBAddress string,      // [R]
	outDBDataBase string,     // [R]
	outDBAuthIsRequired bool, // [R]
	outDBAuthDataBase string, // [R]
	outDBUsername string,     // [R]
	outDBPassword string,     // [R]
) error {

	var err error

	// Database Information Message.
	fmt.Printf("Database Address: '%v'.\r\n", outDBAddress)
	fmt.Printf("Database Base: '%v'.\r\n", outDBDataBase)

	if outDBAuthIsRequired {
		fmt.Printf(
			"Database Authentication Base: '%v'.\r\n",
			outDBAuthDataBase,
		)
		fmt.Printf("Database Username: '%v'.\r\n", outDBUsername)
		fmt.Printf("Database Password: '%v'.\r\n", "***")
	}

	// Connect to Database.
	if !outDBAuthIsRequired {

		// Simple Connection without Authentication.
		dbSession, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Addrs: []string{
					outDBAddress,
				},
				Database: outDBDataBase,
			},
		)
		if err != nil {
			return err
		}

	} else {

		// Connection with Authentication.
		dbSession, err = mgo.DialWithInfo(
			&mgo.DialInfo{
				Addrs: []string{
					outDBAddress,
				},
				Database: outDBDataBase,

				// Authentication Data.
				Source:   outDBAuthDataBase, // = Auth. DB.
				Password: outDBPassword,
				Username: outDBUsername,
			},
		)
		if err != nil {
			return err
		}
	}

	// Save internal Parameters.
	dbDatabase = outputDBDataBase
	dbCollection = dbSession.DB(dbDatabase).C(MongoDBCollectionName)

	return nil
}

// Inserts an Object's Data into the Database.
func planetInsertIntoDB(object Planet) error {

	var err error

	// Insert Data.
	err = dbCollection.Insert(object)
	if err != nil {
		return err
	}

	return nil
}
