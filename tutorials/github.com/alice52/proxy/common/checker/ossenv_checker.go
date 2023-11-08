package checker

import (
	"fmt"
	"os"
)

func CheckOssEnv() (string, string, error) {
	accessID := os.Getenv("OSS_ACCESS_KEY_ID")
	if accessID == "" {
		return "", "", fmt.Errorf("access key id is empty!")
	}
	accessKey := os.Getenv("OSS_ACCESS_KEY_SECRET")
	if accessKey == "" {
		return "", "", fmt.Errorf("access key secret is empty!")
	}

	return accessID, accessKey, nil
}
