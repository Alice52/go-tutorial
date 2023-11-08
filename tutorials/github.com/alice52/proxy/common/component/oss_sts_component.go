package component

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/alice52/proxy/common/constants"
	"time"
)

func CreateClient(accessKeyId *string, accessKeySecret *string) (*sts.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        tea.String("sts." + constants.Region),
	}

	return sts.NewClient(config)
}

func GetAssumeRole(ak, sk, arn string) (*sts.AssumeRoleResponse, error) {
	client, err := CreateClient(tea.String(ak), tea.String(sk))
	if err != nil {
		return nil, err
	}

	policy := `{
			"Version": "1",
			"Statement": [
				{
					"Effect": "Allow",
					"Action": "oss:GetObject",
					"Resource": "acs:oss:*:*:project-ec/*"
				}
			]
		}`
	assumeRoleRequest := &sts.AssumeRoleRequest{
		RoleArn:         tea.String(arn),
		Policy:          &policy,
		DurationSeconds: tea.Int64(int64(time.Hour.Seconds())),
		RoleSessionName: tea.String("sts"),
	}

	return client.AssumeRoleWithOptions(assumeRoleRequest, &util.RuntimeOptions{})
}
