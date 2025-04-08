package s3storage

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
)

func DownloadFileToMemory(key string) (*bytes.Buffer, error) {
	conn := s3.New(storage.session)

	output, err := conn.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(storage.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeNoSuchKey {
			return nil, ErrFileExists
		}
		return nil, err
	}
	defer output.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, output.Body); err != nil {
		return nil, err
	}

	return buf, nil
}
