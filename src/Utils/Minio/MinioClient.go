package MinioClient

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func ConnectMinio() (*minio.Client, error) {
	endpoint := "194.233.95.186:9199"
	accessKeyID := "admin"
	secretAccessKey := "RveWRG9uyfTgw57r"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
