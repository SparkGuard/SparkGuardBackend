package s3

import (
	"github.com/aws/aws-sdk-go/service/s3"
)

func (storage *S3Storage) isBucketExists(bucket string) (exists bool) {
	conn := s3.New(storage.session)

	_, err := conn.HeadBucket(&s3.HeadBucketInput{
		Bucket: &bucket,
	})

	return err == nil
}

func (storage *S3Storage) IsFileExists(filePath string) (exists bool) {
	conn := s3.New(storage.session)

	_, err := conn.HeadObject(&s3.HeadObjectInput{
		Bucket: &storage.bucket,
		Key:    &filePath,
	})

	return err == nil
}
