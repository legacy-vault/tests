// test_a.go.

// Minio Functions.

/*
	Introduction.

	Minio is a cloud storage server released under Apache License v2,
	compatible with Amazon S3. 'Minio Inc' is the prime developer of the
	Minio cloud storage stack. 'Minio Inc' is a Silicon valley based technology
	startup, founded by Anand Babu Periasamy (AB), Garima Kapoor and
	Harshavardhana in November, 2014.

	Source: https://en.wikipedia.org/wiki/Minio.

*/

// Date: 2018-09-17.

package minio

import (
	"github.com/minio/minio-go"
	"log"
)

func TestA() error {

	var cfg MinioConfiguration
	var err error
	var minioClient minio.Client

	cfg.Endpoint = "play.minio.io:9000"
	cfg.AccessKeyID = "Q3AM3UQ867SPQQA43P2F"
	cfg.SecretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	cfg.SSLisUsed = true
	cfg.BucketName = "xxx-xxx"
	cfg.BucketLocation = "us-east-1"

	// 1. Initialize Minio Errors.
	ErrorsInitialize()

	// 2. Initialize Minio Client Connection.
	err = ClientInitialize(cfg, &minioClient)
	if err != nil {
		return err
	}

	// Minio Database provides only a "Constructor" as it is said in its
	// Documentation. I.e. there is no 'Close' or 'Disconnect' Method provided.
	// Source: https://docs.minio.io/docs/golang-client-api-reference.

	// Log.
	log.Println("Minio Initialization: [OK]")
	log.Printf("Minio Client: %+v.\r\n", minioClient) //!

	// 3. Initialize Minio Bucket.
	err = BucketInitialize(&minioClient, cfg)
	if err != nil {
		// Skip the 'Bucket is not empty' Error.
		if err.Error() != ErrBucketNotEmpty.Error() {
			return err
		}
	}

	// 4. Upload a Test File to the Bucket.
	err = FileUploadA(&minioClient, cfg)
	if err != nil {
		return err
	}

	// 5. Download a Test File to the Bucket.
	err = FileDownloadA(&minioClient, cfg)
	if err != nil {
		return err
	}

	// 6. Delete the Test File from the Bucket.
	err = FileDeleteA(&minioClient, cfg)
	if err != nil {
		return err
	}

	// 7. Test Copy.
	err = CopyA(&minioClient, cfg)
	if err != nil {
		return err
	}

	// 8. Test Partial Upload, Join and Download.
	err = PartialUploadAndJoinA(&minioClient, cfg)
	if err != nil {
		return err
	}

	return nil
}
