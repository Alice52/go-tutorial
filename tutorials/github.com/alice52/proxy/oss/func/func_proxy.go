package _func

import (
	"context"
	"encoding/json"
	"github.com/alice52/proxy/common/checker"
	"github.com/alice52/proxy/common/component"
	"github.com/alice52/proxy/common/log"
	"github.com/alice52/proxy/common/resp"
	"github.com/alice52/proxy/oss/model"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
)

// HandleHttpRequest check env, params and signature temporary image url for  by oss bucket
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	log.LogFcContext(ctx, req)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 1. check oss env and request params, and respond error if invalid
	ak, sk, objectName, success := CheckOrRespondError(w, req)
	if !success {
		return nil
	}

	// 2. build oss bucket instance
	bucket, err := component.BuildOssBucket(ak, sk)
	if err != nil && resp.RespondError(w, err) {
		return nil
	}

	// 3. do signature for temporary
	signedUrl, err := bucket.SignURL(objectName, oss.HTTPGet, 5*60)
	if err != nil && resp.RespondError(w, err) {
		return nil
	}

	return json.NewEncoder(w).Encode(&model.SignImage{
		Url: signedUrl,
	})
}

func CheckOrRespondError(w http.ResponseWriter, req *http.Request) (string, string, string, bool) {
	// 1. check oss env
	ak, sk, err := checker.CheckOssEnv()
	if err != nil && resp.RespondError(w, err) {
		return "", "", "", false
	}

	// 2. check req params
	objectName, err := checker.CheckReqAndParseObjectName(req)
	if err != nil && resp.RespondError(w, err) {
		return "", "", "", false
	}

	return ak, sk, objectName, true
}
