// file.go.

// Minio File Functions.

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
	"fmt"
	"github.com/minio/minio-go"
	"log"
)

func FileUploadA(mc *minio.Client, mcfg MinioConfiguration) error {

	var err error
	var filePath string
	var objectName string
	var options minio.PutObjectOptions
	var size int64

	// Upload a File: Part I.
	objectName = "My File #1"
	filePath = "./minio/files/1.txt.zip"
	options.ContentType = "application/zip"
	options.NumThreads = 1

	// Upload a File: Part II.
	size, err = mc.FPutObject(
		mcfg.BucketName,
		objectName,
		filePath,
		options,
	)
	if err != nil {
		return err
	}

	// Log.
	log.Println("Uploaded a File:", size, "Bytes.") //!

	return nil
}

func FileDownloadA(mc *minio.Client, mcfg MinioConfiguration) error {

	var err error
	var filePathDownloaded string
	var objectName string
	var options minio.GetObjectOptions

	// Download a File: Part I.
	objectName = "My File #1"
	filePathDownloaded = "./minio/files/1_dl.txt.zip"

	// Download a File: Part II.
	err = mc.FGetObject(
		mcfg.BucketName,
		objectName,
		filePathDownloaded,
		options,
	)
	if err != nil {
		return err
	}

	// Log.
	log.Println("Downloaded a File:", filePathDownloaded, ".") //!

	return nil
}

func FileDeleteA(mc *minio.Client, mcfg MinioConfiguration) error {

	var err error
	var objectName string

	// Remove a File Object.
	objectName = "My File #1"

	// Remove a File Object.
	err = mc.RemoveObject(
		mcfg.BucketName,
		objectName,
	)
	if err != nil {
		return err
	}

	// Log.
	log.Println("Removed an Object:", objectName, ".") //!

	return nil
}

func CopyA(mc *minio.Client, mcfg MinioConfiguration) error {

	var doneCh chan struct{}
	var dst minio.DestinationInfo
	var err error
	var list <-chan minio.ObjectInfo
	var listObject minio.ObjectInfo
	var putOptions minio.PutObjectOptions
	var src minio.SourceInfo

	// 1. Upload a File.
	putOptions.ContentType = "text/plain"
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
	putOptions.NumThreads = 1
	//
	_, err = mc.FPutObject(
		mcfg.BucketName,
		"Copy Source",
		"./minio/files/copy_source.txt",
		putOptions,
	)
	if err != nil {
		return err
	}
	// Prepare a Minio Source.
	src = minio.NewSourceInfo(
		mcfg.BucketName,
		"Copy Source",
		nil,
	)

	// 2. Copy the File.
	dst, err = minio.NewDestinationInfo(
		mcfg.BucketName,
		"Copy Destination",
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	err = mc.CopyObject(dst, src)
	if err != nil {
		return err
	}

	// 3. List Objects.
	doneCh = make(chan struct{})
	defer close(doneCh)
	list = mc.ListObjects(
		mcfg.BucketName,
		"",
		false,
		doneCh,
	)
	// Output to Console.
	fmt.Print("LIST: ")
	for listObject = range list {
		if listObject.Err != nil {
			return listObject.Err
		}
		fmt.Print(listObject, " ")
	}
	fmt.Println("")

	// 4. Clean-up.
	err = mc.RemoveObject(
		mcfg.BucketName,
		"Copy Destination",
	)
	if err != nil {
		return err
	}
	err = mc.RemoveObject(
		mcfg.BucketName,
		"Copy Source",
	)
	if err != nil {
		return err
	}

	// 5. List Objects again.
	doneCh = make(chan struct{})
	defer close(doneCh)
	list = mc.ListObjects(
		mcfg.BucketName,
		"",
		false,
		doneCh,
	)
	// Output to Console.
	fmt.Print("LIST: ")
	for listObject = range list {
		if listObject.Err != nil {
			return listObject.Err
		}
		fmt.Print(listObject, " ")
	}
	fmt.Println("")

	return nil
}

func PartialUploadAndJoinA(mc *minio.Client, mcfg MinioConfiguration) error {

	const PartsCount = 2

	var destination minio.DestinationInfo
	var err error
	var getOptions minio.GetObjectOptions
	var putOptions minio.PutObjectOptions
	var sources []minio.SourceInfo

	// 1. Upload Two Parts.
	sources = make([]minio.SourceInfo, PartsCount)
	putOptions.ContentType = "text/plain"
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
	putOptions.NumThreads = 1

	// 1.1. Upload the first File.
	_, err = mc.FPutObject(
		mcfg.BucketName,
		"Part #1",
		"./minio/files/file.part_1.txt",
		putOptions,
	)
	if err != nil {
		return err
	}
	// Prepare a Minio Source.
	sources[0] = minio.NewSourceInfo(
		mcfg.BucketName,
		"Part #1",
		nil,
	)

	// 1.2. Upload the second File.
	_, err = mc.FPutObject(
		mcfg.BucketName,
		"Part #2",
		"./minio/files/file.part_2.txt",
		putOptions,
	)
	if err != nil {
		return err
	}
	// Prepare a Minio Source.
	sources[1] = minio.NewSourceInfo(
		mcfg.BucketName,
		"Part #2",
		nil,
	)

	// 2. Join Files into a single File.
	destination, err = minio.NewDestinationInfo(
		mcfg.BucketName,
		"All Parts",
		nil,
		nil,
	)
	if err != nil {
		return err
	}
	err = mc.ComposeObject(destination, sources)
	if err != nil {
		return err
	}

	// 3. Download the joined File.
	err = mc.FGetObject(
		mcfg.BucketName,
		"All Parts",
		"./minio/files/file.all_parts.txt",
		getOptions,
	)
	if err != nil {
		return err
	}

	// 4. Remove all Parts and the joined File.
	err = mc.RemoveObject(
		mcfg.BucketName,
		"Part #1",
	)
	if err != nil {
		return err
	}
	err = mc.RemoveObject(
		mcfg.BucketName,
		"Part #2",
	)
	if err != nil {
		return err
	}
	err = mc.RemoveObject(
		mcfg.BucketName,
		"All Parts",
	)
	if err != nil {
		return err
	}

	// Hooray! Job is done!
	// Log.
	log.Printf(
		"Downloaded a joined File to: '%v'.\r\n",
		"./minio/files/file.all_parts.txt",
	) //!

	return nil
}
