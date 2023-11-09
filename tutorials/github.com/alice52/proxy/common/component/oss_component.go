package component

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/alice52/proxy/common/constants"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Deprecated: BuildOssStsClient
func BuildOssStsClient(ak, sk, arn string) (*oss.Client, error) {

	if ak == "" || sk == "" || arn == "" {
		return nil, fmt.Errorf("ak or sk or arn is empty")
	}

	role, err := GetAssumeRole(ak, sk, arn)
	if err != nil {
		return nil, fmt.Errorf("error to get assueme role, due to %s", err.Error())
	}

	provider := stsCredentialsProvider{
		&stsCredentials{
			AccessKeyId:     tea.StringValue(role.Body.Credentials.AccessKeyId),
			AccessKeySecret: tea.StringValue(role.Body.Credentials.AccessKeySecret),
			SecurityToken:   tea.StringValue(role.Body.Credentials.SecurityToken),
		},
	}

	return oss.New(constants.Endpoint, ak, sk, oss.SetCredentialsProvider(&provider))
}

// Deprecated: BuildOssStsBucket build oss bucket of sts
func BuildOssStsBucket(ak, sk, arn string) (*oss.Bucket, error) {
	// 1. build oss client
	client, err := BuildOssStsClient(ak, sk, arn)
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
