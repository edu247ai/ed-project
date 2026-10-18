package s3

import (
	"bytes"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Adapter struct {
	client *s3.S3
	bucket string
}

func New(client *s3.S3, bucket string) *S3Adapter {
	return &S3Adapter{client: client, bucket: bucket}
}

func (s *S3Adapter) Upload(cid string, data []byte) error {
	_, err := s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(cid),
		Body:   bytes.NewReader(data),
	})
	return err
}

func (s *S3Adapter) Download(cid string) ([]byte, error) {
	out, err := s.client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(cid),
	})
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()

	return io.ReadAll(out.Body)
}
