package component

import (
	"fmt"
	"github.com/alice52/proxy/common/constants"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Deprecated: BuildOssRawClient
func BuildOssRawClient(ak, sk, securityToken string) (*oss.Client, error) {

	if ak == "" || sk == "" || securityToken == "" {
		return nil, fmt.Errorf("ak or sk or securityToken is empty!")
	}

	provider := stsCredentialsProvider{
		&stsCredentials{
			AccessKeyId:     ak,
			AccessKeySecret: sk,
			SecurityToken:   securityToken,
		},
	}

	return oss.New(constants.Endpoint, ak, sk, oss.SetCredentialsProvider(&provider))
}

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
