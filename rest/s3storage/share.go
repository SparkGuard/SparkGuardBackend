package s3

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

func (storage *S3Storage) ShareFile(key string) (url string, err error) {
	conn := s3.New(storage.session)

	req, _ := conn.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &storage.bucket,
		Key:    &key,
	})

	return req.Presign(time.Hour)
}
