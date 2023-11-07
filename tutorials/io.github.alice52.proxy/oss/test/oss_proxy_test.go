package test

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io.github.alice52.proxy/oss/constants"
	"os"
	"testing"
)

func TestOssProxy(t *testing.T) {
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	client, err := oss.New(constants.Endpoint, os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"), oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	bucketName := constants.BucketName
	objectName := "46.png"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 60)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Printf("Sign Url:%s\n", signedURL)
}
