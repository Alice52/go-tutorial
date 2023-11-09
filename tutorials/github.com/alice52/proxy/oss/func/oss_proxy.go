package _func

import (
	"context"
	"encoding/json"
	"github.com/alice52/proxy/common/checker"
	"github.com/alice52/proxy/common/constants"
	"github.com/alice52/proxy/common/model"
	_oss "github.com/alice52/proxy/common/oss"
	"github.com/alice52/proxy/common/util"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
)

// HandleHttpRequest check env, params and signature temporary image url for  by oss bucket
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	util.LogFcContext(ctx, req)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 1. check oss env and request params, and respond error if invalid
	ak, sk, objectName, success := CheckOrRespondError(w, req)
	if !success {
		return nil
	}

	// 2. build oss bucket instance
	bucket, err := _oss.BuildOssBucket(ak, sk)
	if err != nil && util.RespondError(constants.InternalError, w, err) {
		return nil
	}

	// 3. do signature for temporary
	signedUrl, err := bucket.SignURL(objectName, oss.HTTPGet, constants.ExpireSeconds)
	if err != nil && util.RespondError(constants.InternalError, w, err) {
		return nil
	}

	return json.NewEncoder(w).Encode(&model.R{
		ErrCode: constants.Success,
		Url:     signedUrl,
	})
}

func CheckOrRespondError(w http.ResponseWriter, req *http.Request) (string, string, string, bool) {
	// 1. check oss env
	ak, sk, err := checker.CheckOssEnv()
	if err != nil && util.RespondError(constants.InternalError, w, err) {
		return "", "", "", false
	}

	// 2. check req params
	objectName, err := checker.CheckReqAndParseObjectName(req)
	if err != nil && util.RespondError(constants.BadRequest, w, err) {
		return "", "", "", false
	}

	return ak, sk, objectName, true
}
