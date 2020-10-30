package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

var Client *minio.Client

type MinioConfig struct {
	Host           string
	Port           string
	AccessKey      string
	SecretKey      string
	MainBucketName string
}

func StartMinioConnection(cfg *MinioConfig) {
	endPoint := cfg.Host + ":" + cfg.Port
	client, err := minio.New(endPoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
	})
	if err != nil {
		log.Err(err).Msg("Can not create minio client.")
		panic(err)
	}

	Client = client

	if !ExistBucket(cfg.MainBucketName) {
		err := CreateBucket(cfg.MainBucketName)
		if err != nil {
			panic(err)
		}
	}
}

func BuildMinioConfig() *MinioConfig {
	config := MinioConfig{
		Host:           "localhost",
		Port:           "9000",
		AccessKey:      "AKIAIOSFODNN7EXAMPLE",
		SecretKey:      "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		MainBucketName: "file-service",
	}

	return &config
}
