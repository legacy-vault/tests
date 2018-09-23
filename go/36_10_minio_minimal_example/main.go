// main.go.

// Minio Database Test.
// Minimal Example for 'stackoverflow.com' Community.

// Aim:
// Two Files must be concatenated inside MINIO DB and downloaded by the Client.
// Instead, the 'ComposeObject' Function returns an Error:
// "At least one of the pre-conditions you specified did not hold".

// Date: 2018-09-18.

package main

import (
	"github.com/minio/minio-go"
	"log"
	"os"
)

const FilePath1 = "./minio/files/file.part_1.txt"
const FilePath2 = "./minio/files/file.part_2.txt"
const FilePath1and2 = "./minio/files/file.all_parts.txt"

const ObjContentType = "text/plain"

const ObjName1 = "Part #1"
const ObjName2 = "Part #2"
const ObjName1and2 = "All Parts"

const PartsCount = 2

type MinioConfiguration struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	SSLisUsed       bool
	BucketName      string
	BucketLocation  string
}

func main() {

	var bucketExists bool
	var cfg MinioConfiguration
	var destination minio.DestinationInfo
	var err error
	var getOptions minio.GetObjectOptions
	var minioClient *minio.Client
	var putOptions minio.PutObjectOptions
	var sources []minio.SourceInfo

	// 1. Configure Minio Client. This Data is public, for Tests.
	cfg.Endpoint = "play.minio.io:9000"
	cfg.AccessKeyID = "Q3AM3UQ867SPQQA43P2F"
	cfg.SecretAccessKey = "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	cfg.SSLisUsed = true
	cfg.BucketName = "stackoverflow-com-minimal-example"
	cfg.BucketLocation = "us-east-1"

	// 2. Create a Client, i.e. "connect" to Minio Database.
	minioClient, err = minio.New(
		cfg.Endpoint,
		cfg.AccessKeyID,
		cfg.SecretAccessKey,
		cfg.SSLisUsed,
	)

	// 3. Create a Bucket if it does not exist.
	bucketExists, err = minioClient.BucketExists(cfg.BucketName)
	check_error(err)
	if (bucketExists == false) {
		err = minioClient.MakeBucket(cfg.BucketName, cfg.BucketLocation)
		check_error(err)
	}

	// 4. Upload Two Parts.
	sources = make([]minio.SourceInfo, PartsCount)
	putOptions.ContentType = ObjContentType
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
	putOptions.NumThreads = 1

	// 4.1. Upload the first File.
	_, err = minioClient.FPutObject(
		cfg.BucketName,
		ObjName1,
		FilePath1,
		putOptions,
	)
	check_error(err)
	// Prepare a Minio Source.
	sources[0] = minio.NewSourceInfo(
		cfg.BucketName,
		ObjName1,
		nil,
	)

	// 4.2. Upload the second File.
	_, err = minioClient.FPutObject(
		cfg.BucketName,
		ObjName2,
		FilePath2,
		putOptions,
	)
	check_error(err)
	// Prepare a Minio Source.
	sources[1] = minio.NewSourceInfo(
		cfg.BucketName,
		ObjName2,
		nil,
	)

	// 5. Join Files into a single File.
	destination, err = minio.NewDestinationInfo(
		cfg.BucketName,
		ObjName1and2,
		nil,
		nil,
	)
	check_error(err)
	err = minioClient.ComposeObject(destination, sources)
	check_error(err)

	// 6. Download the joined File.
	err = minioClient.FGetObject(
		cfg.BucketName,
		ObjName1and2,
		FilePath1and2,
		getOptions,
	)
	check_error(err)

	// 7. Remove all Parts and the joined File.
	err = minioClient.RemoveObject(
		cfg.BucketName,
		ObjName1,
	)
	check_error(err)
	err = minioClient.RemoveObject(
		cfg.BucketName,
		ObjName2,
	)
	check_error(err)
	err = minioClient.RemoveObject(
		cfg.BucketName,
		ObjName1and2,
	)
	check_error(err)

	// Hooray! Job is done!
	// Log.
	log.Printf(
		"Downloaded a joined File to: '%v'.\r\n",
		FilePath1and2,
	) //!
}

func check_error(e error) {

	if e != nil {
		log.Println(e.Error())
		os.Exit(1)
	}
}
