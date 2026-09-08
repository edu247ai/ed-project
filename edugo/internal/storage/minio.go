package storage

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioStorage struct {
	client     *minio.Client
	bucketName string
}

// ------------------------------------------------------------
// CONSTRUCTOR
// ------------------------------------------------------------
func NewMinio(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinioStorage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	// Ensure bucket exists
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}
	}

	return &MinioStorage{
		client:     client,
		bucketName: bucket,
	}, nil
}

// ------------------------------------------------------------
// IMPLEMENT STORAGE INTERFACE
// ------------------------------------------------------------

func (m *MinioStorage) Save(path string, data io.Reader) error {
	_, err := m.client.PutObject(
		context.Background(),
		m.bucketName,
		path,
		data,
		-1,
		minio.PutObjectOptions{},
	)
	return err
}

func (m *MinioStorage) Get(path string) (io.ReadCloser, error) {
	return m.client.GetObject(
		context.Background(),
		m.bucketName,
		path,
		minio.GetObjectOptions{},
	)
}

func (m *MinioStorage) Delete(path string) error {
	return m.client.RemoveObject(
		context.Background(),
		m.bucketName,
		path,
		minio.RemoveObjectOptions{},
	)
}

func (m *MinioStorage) PresignURL(path string, expiry time.Duration) (string, error) {
	reqParams := make(url.Values)
	u, err := m.client.PresignedGetObject(
		context.Background(),
		m.bucketName,
		path,
		expiry,
		reqParams,
	)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
