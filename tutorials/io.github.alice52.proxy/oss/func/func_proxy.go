package _func

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
)

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	ak, _ := base64.StdEncoding.DecodeString("TFRBSTV0TVVXWlhnS3M4NjhnUUs1cVFk")
	sk, _ := base64.StdEncoding.DecodeString("NnUzNFA2OXd4S1g5dHhBOHRqaG1lR2phUzhOV1ZY")

	client, err := oss.New("https://oss-cn-hangzhou.aliyuncs.com", string(ak), string(sk), oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucketName := "project-ec"
	objectName := "45+.png"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, 600)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Printf("Sign Url:%s\n", signedURL)
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(signedURL))

	return nil
}
