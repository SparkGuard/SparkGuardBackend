package s3storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func isBucketExists(bucket string) (exists bool) {
	conn := s3.New(storage.session)

	_, err := conn.HeadBucket(&s3.HeadBucketInput{
		Bucket: &bucket,
	})

	return err == nil
}

func IsFileExists(filePath string) (exists bool) {
	conn := s3.New(storage.session)

	_, err := conn.HeadObject(&s3.HeadObjectInput{
		Bucket: &storage.bucket,
		Key:    &filePath,
	})

	return err == nil
}
