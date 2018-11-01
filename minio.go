package faasflowMinioStateManager

import (
	"fmt"
	minio "github.com/minio/minio-go"
	faasflow "github.com/s8sg/faasflow"
)

type MinioStateManager struct {
	initialized bool
	minioClient *minio.Client
}

func GetMinioStateManager() (*MinioStateManager, error) {

	minioStateManager := &MinioStateManager{}

	region := regionName()
	bucketName := bucketName()

	minioClient, connectErr := connectToMinio(region)
	if connectErr != nil {
		return nil, fmt.Errorf("Failed to initialize minio, error %s", connectErr.Error())
	}

	minioStateManager.initialized = true

	return minioStateManager, nil
}

func (minioState *MinioStateManager) Set(key string, value interface{}) error {

	return nil
}

func (minioState *MinioStateManager) Get(key string) (interface{}, error) {

	return nil
}

func (minioState *MinioStateManager) Del(key string) error {

	return nil
}

func connectToMinio(region string) (*minio.Client, error) {

	endpoint := os.Getenv("s3_url")

	tlsEnabled := tlsEnabled()

	secretKey, _ := sdk.ReadSecret("s3-secret-key")
	accessKey, _ := sdk.ReadSecret("s3-access-key")

	return minio.New(endpoint, accessKey, secretKey, tlsEnabled)
}

// getPath produces a string such as pipeline/
func getPath(bucket string, p *sdk.PipelineLog) string {
	fileName := "faasflow.state"
	return fmt.Sprintf("%s/%s", bucket, fileName)
}

func tlsEnabled() bool {
	if connection := os.Getenv("s3_tls"); connection == "true" || connection == "1" {
		return true
	}
	return false
}

func bucketName() string {
	bucketName, exist := os.LookupEnv("s3_bucket")
	if exist == false || len(bucketName) == 0 {
		bucketName = "pipeline"
		log.Printf("Bucket name not found, set to default: %v\n", bucketName)
	}
	return bucketName
}

func regionName() string {
	regionName, exist := os.LookupEnv("s3_region")
	if exist == false || len(regionName) == 0 {
		regionName = "us-east-1"
	}
	return regionName
}
