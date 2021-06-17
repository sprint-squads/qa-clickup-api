package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sprint-squads/qa-clickup-api/pkg/config"
	"mime/multipart"
)

func Get(cfg *config.Config) *minio.Client {
	endpoint := cfg.Minio.Url
	accessKeyID := cfg.Minio.AccessKey
	secretAccessKey := cfg.Minio.SecretKey
	useSSL := cfg.Minio.UseSSL

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		fmt.Println(err)
		panic(1)
	}

	return client
}

func UploadToMinio(minioClient *minio.Client, bucketName string, objectName string, file *multipart.FileHeader, contentType string) (uploadUrl string, err error) {
	var src multipart.File
	src, err = file.Open()
	if err != nil {
		return
	}
	defer src.Close()

	var uploadInfo minio.UploadInfo
	uploadInfo, err = minioClient.PutObject(context.Background(), bucketName, objectName, src,-1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return
	}

	uploadUrl = uploadInfo.Key
	return
}
