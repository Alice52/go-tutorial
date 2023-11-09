package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//type Credentials interface {
//	GetAccessKeyID() string
//	GetAccessKeySecret() string
//	GetSecurityToken() string
//}

type stsCredentials struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

func (credentials *stsCredentials) GetAccessKeyID() string {
	return credentials.AccessKeyId
}

func (credentials *stsCredentials) GetAccessKeySecret() string {
	return credentials.AccessKeySecret
}

func (credentials *stsCredentials) GetSecurityToken() string {
	return credentials.SecurityToken
}

//type CredentialsProvider interface {
//	GetCredentials() Credentials
//}

type stsCredentialsProvider struct {
	cred oss.Credentials
}

func (credentials *stsCredentialsProvider) GetCredentials() oss.Credentials {

	return &stsCredentials{
		AccessKeyId:     credentials.cred.GetAccessKeyID(),
		AccessKeySecret: credentials.cred.GetAccessKeySecret(),
		SecurityToken:   credentials.cred.GetSecurityToken(),
	}
}
