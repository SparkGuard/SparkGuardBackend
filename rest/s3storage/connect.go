package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Connect(endpoint, region, bucket string) (storage *S3Storage, err error) {
	storage = new(S3Storage)

	storage.session, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    &endpoint,
		Region:      &region,
	})

	storage.bucket = bucket

	if !storage.isBucketExists(bucket) {
		err = ErrBucketNotExists
		return
	}

	return
}
