// minio.go.

// General Minio Data.

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
)

type MinioConfiguration struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	SSLisUsed       bool
	BucketName      string
	BucketLocation  string
}

// Unfortunately, most of Errors are not public in 'minio' Library.
// So, we have to create them by ourselves.
const ErrMsgBucketNotEmpty = "The bucket you tried to delete is not empty"

var ErrBucketNotEmpty error

func ErrorsInitialize() {

	ErrBucketNotEmpty = errors.New(ErrMsgBucketNotEmpty)
}

func ClientInitialize(mcfg MinioConfiguration, mc *minio.Client) error {

	var err error
	var mcTmp *minio.Client

	// Create a Client.
	mcTmp, err = minio.New(
		mcfg.Endpoint,
		mcfg.AccessKeyID,
		mcfg.SecretAccessKey,
		mcfg.SSLisUsed,
	)
	if err != nil {
		return err
	}

	*mc = *mcTmp
	return nil
}
