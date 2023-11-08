package component

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io.github.alice52.proxy/oss/constants"
)

// BuildOssClient build oss client
func BuildOssClient(ak, sk string) (*oss.Client, error) {
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		return nil, err
	}

	return oss.New(constants.Endpoint, ak, sk, oss.SetCredentialsProvider(&provider))
}

// BuildOssBucket build oss bucket
func BuildOssBucket(ak string, sk string) (*oss.Bucket, error) {
	// 1. build oss client
	client, err := BuildOssClient(ak, sk)
	if err != nil {
		return nil, err
	}

	// 2. build oss bucket
	bucket, err := client.Bucket(constants.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
