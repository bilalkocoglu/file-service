package minio

import (
	_const "github.com/imminoglobulin/file-service/pkg/const"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
	"os"
)

var Client *minio.Client

type MinioConfig struct {
	URL            string
	AccessKey      string
	SecretKey      string
	MainBucketName string
}

func StartMinioConnection(cfg *MinioConfig) {
	client, err := minio.New(cfg.URL, &minio.Options{
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
		URL:            _const.DefaultMinioUrl,
		AccessKey:      _const.DefaultMinioAccessKey,
		SecretKey:      _const.DefaultMinioSecretKey,
		MainBucketName: _const.DefaultMainBucketName,
	}

	minioUrl := os.Getenv("MINIO_URL")
	if minioUrl != "" {
		config.URL = minioUrl
	}

	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	if accessKey != "" {
		config.AccessKey = accessKey
	}

	secretKey := os.Getenv("MINIO_SECRET_KEY")
	if secretKey != "" {
		config.SecretKey = secretKey
	}

	mainBucketName := os.Getenv("MAIN_BUCKET_NAME")
	if mainBucketName != "" {
		config.MainBucketName = mainBucketName
	}

	return &config
}
