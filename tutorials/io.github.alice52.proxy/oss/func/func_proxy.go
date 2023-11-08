package _func

import (
	"context"
	"encoding/json"
	"io.github.alice52.proxy/oss/component"
	"io.github.alice52.proxy/oss/constants"
	"io.github.alice52.proxy/oss/model"
	"strings"

	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
)

// HandleHttpRequest check env, params and signature temporary image url for  by oss bucket
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 1. check oss env and request params
	ak, sk, objectName, success := checkEnvAndReq(w, req)
	if !success {
		return nil
	}

	// 2. build oss bucket instance
	bucket, err := component.BuildOssBucket(ak, sk)
	if err != nil && response(w, err) {
		return nil
	}

	// 3. do signature for temporary
	signedUrl, err := bucket.SignURL(objectName, oss.HTTPGet, 5*60)
	if err != nil && response(w, err) {
		return nil
	}

	return json.NewEncoder(w).Encode(&model.SignImage{
		Url: signedUrl,
	})
}

func checkEnvAndReq(w http.ResponseWriter, req *http.Request) (string, string, string, bool) {
	// 1. check oss env
	ak, sk, err := checkOssEnv()
	if err != nil && response(w, err) {
		return "", "", "", false
	}

	// 2. check req params
	objectName, err := checkReqAndParseObjectName(req)
	if err != nil && response(w, err) {
		return "", "", "", false
	}

	return ak, sk, objectName, true
}

func response(w http.ResponseWriter, err error) bool {
	err = json.NewEncoder(w).Encode(&model.ErrorResult{
		Msg: err.Error(),
	})

	if err != nil {
		fmt.Println("Error responding error response")
	}

	return true
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
