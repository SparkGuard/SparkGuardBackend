package s3

import (
	"bytes"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

func (storage *S3Storage) UploadFile(filePath string, file io.Reader) (err error) {
	var buf bytes.Buffer

	if _, err = io.Copy(&buf, file); err != nil {
		return
	}

	conn := s3.New(storage.session)

	_, err = conn.PutObject(&s3.PutObjectInput{
		Bucket: &storage.bucket,
		Key:    &filePath,
		Body:   bytes.NewReader(buf.Bytes()),
	})

	return
}

func (storage *S3Storage) UploadFileSafe(filePath string, file io.Reader) (err error) {
	if storage.IsFileExists(filePath) {
		return ErrFileExists
	}

	return storage.UploadFile(filePath, file)
}
