package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

func ExistBucket(bucketName string) bool {
	ctx := context.Background()
	exists, err := Client.BucketExists(ctx, bucketName)
	if err != nil {
		panic(err)
	}
	return exists
}

func CreateBucket(bucketName string) error {
	ctx := context.Background()
	return Client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
}

func UploadFile(bucketName string, fileName string, file io.Reader) (minio.UploadInfo, error) {
	ctx := context.Background()
	return Client.PutObject(ctx, bucketName, fileName, file, 324, minio.PutObjectOptions{})
}
