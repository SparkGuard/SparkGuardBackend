package s3storage

import "github.com/aws/aws-sdk-go/aws/session"

type S3Storage struct {
	session *session.Session
	bucket  string
}

type S3Error string

func (err S3Error) Error() string {
	return string(err)
}

const ErrBucketNotExists = S3Error("bucket not exists")
const ErrFileExists = S3Error("file already exists")
