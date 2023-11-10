package test

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/alice52/proxy/common/constants"
	_oss "github.com/alice52/proxy/common/oss"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

func TestOssProxy(t *testing.T) {
	rak := os.Getenv("OSS_ACCESS_KEY_ID")
	rsk := os.Getenv("OSS_ACCESS_KEY_SECRET")
	arn := os.Getenv("OSS_ARN_KEY")
	role, err := _oss.GetAssumeRole(rak, rsk, arn)
	if err != nil {
		fmt.Printf("%v", err)
	}

	ak := tea.StringValue(role.Body.Credentials.AccessKeyId)
	sk := tea.StringValue(role.Body.Credentials.AccessKeySecret)
	st := tea.StringValue(role.Body.Credentials.SecurityToken)

	client, err := _oss.BuildOssStsClient(ak, sk, st)
	if err != nil {
		return
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
