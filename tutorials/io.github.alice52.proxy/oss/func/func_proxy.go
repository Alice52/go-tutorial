package _func

import (
	"context"
	"encoding/json"
	"io.github.alice52.proxy/oss/constants"
	"io.github.alice52.proxy/oss/model"
	"strings"

	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
)

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 1. check oss env
	ak, sk, err := checkOssEnv()
	if err != nil {
		return response(w, err)
	}

	// 2. check req params
	objectName, err := checkReqAndParseObjectName(req)
	if err != nil {
		return response(w, err)
	}

	// 3. build oss client
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	client, err := oss.New(constants.Endpoint, ak, sk, oss.SetCredentialsProvider(&provider))
	if err != nil {
		return response(w, err)
	}

	// 4. build oss bucket
	bucket, err := client.Bucket(constants.BucketName)
	if err != nil {
		return response(w, err)
	}

	// 5. do signature for temporary
	signedUrl, err := bucket.SignURL(objectName, oss.HTTPGet, 5*60)
	if err != nil {
		return response(w, err)
	}

	return json.NewEncoder(w).Encode(&model.SignImage{
		Url: signedUrl,
	})
}

func response(w http.ResponseWriter, err error) error {
	return json.NewEncoder(w).Encode(&model.ErrorResult{
		Msg: err.Error(),
	})
}

func checkReqAndParseObjectName(req *http.Request) (string, error) {

	params := req.URL.Query()
	url := params.Get("url")
	token := params.Get("token")

	if url == "" {
		return "", fmt.Errorf("url is empty!")
	}
	if token == "" {
		return "", fmt.Errorf("token is empty!")
	}

	return strings.Split(url, constants.Endpoint+"/")[1], nil
}

func checkOssEnv() (string, string, error) {
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
