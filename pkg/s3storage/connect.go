package s3storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var storage *S3Storage

func Connect(endpoint, region, bucket string) (err error) {
	fmt.Println("Starting connection to S3...")

	storage = new(S3Storage)

	storage.session, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewEnvCredentials(),
		Endpoint:    &endpoint,
		Region:      &region,
	})

	storage.bucket = bucket

	if !isBucketExists(bucket) {
		err = ErrBucketNotExists
		return
	}

	fmt.Println("Successfully established connection to S3!")
	return
}
