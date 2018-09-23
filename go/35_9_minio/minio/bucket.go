// bucket.go.

// Minio Bucket Functions.

/*
	Introduction.

	Minio is a cloud storage server released under Apache License v2,
	compatible with Amazon S3. 'Minio Inc' is the prime developer of the
	Minio cloud storage stack. 'Minio Inc' is a Silicon valley based technology
	startup, founded by Anand Babu Periasamy (AB), Garima Kapoor and
	Harshavardhana in November, 2014.

	Source: https://en.wikipedia.org/wiki/Minio.

*/

// Date: 2018-09-14.

package minio

import (
	"errors"
	"github.com/minio/minio-go"
	"log"
)

func BucketInitialize(mc *minio.Client, mcfg MinioConfiguration) error {

	var bucketExists bool
	var err error

	// 1. Check Bucket's Existence and delete it if it already exists.
	bucketExists, err = mc.BucketExists(mcfg.BucketName)
	if err != nil {
		return err
	}
	// Log.
	log.Printf(
		"Minio Bucket '%v' Existence: %v.\r\n",
		mcfg.BucketName,
		bucketExists,
	) //!
	if bucketExists {

		// Try do delete an existing Bucket.
		err = mc.RemoveBucket(mcfg.BucketName)
		if err != nil {
			return err
		}

		// Check Removal.
		bucketExists, err = mc.BucketExists(mcfg.BucketName)
		if err != nil {
			return err
		}
		if bucketExists {
			err = errors.New("Existing Bucket has not been deleted")
			return err
		}

		// Log.
		log.Printf(
			"Minio Bucket '%v' has been removed.\r\n",
			mcfg.BucketName,
		) //!
	}

	// 2. No Bucket exists. Create a new Bucket.
	err = mc.MakeBucket(mcfg.BucketName, mcfg.BucketLocation)
	if err != nil {
		return err
	}
	// Check Creation.
	bucketExists, err = mc.BucketExists(mcfg.BucketName)
	if err != nil {
		return err
	}
	// Log.
	log.Printf(
		"Minio Bucket '%v' Existence: %v.\r\n",
		mcfg.BucketName,
		bucketExists,
	) //!
	if !bucketExists {
		err = errors.New("Bucket has not been created")
		return err
	}

	// 3. A new Bucket has been created.
	// Log.
	log.Printf(
		"Minio Bucket '%v' has been created.\r\n",
		mcfg.BucketName,
	) //!

	return nil
}
