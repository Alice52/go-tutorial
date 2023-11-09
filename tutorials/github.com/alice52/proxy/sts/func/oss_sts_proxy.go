package _func

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alice52/proxy/common/constants"
	"github.com/alice52/proxy/common/model"
	_oss "github.com/alice52/proxy/common/oss"
	"github.com/alice52/proxy/common/util"
	_func "github.com/alice52/proxy/oss/func"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
)

// HandleHttpRequest check env, params and signature temporary image url for  by oss bucket
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	util.LogFcContext(ctx, req)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 1. check oss env and request params, and respond error if invalid
	ak, sk, arn, objectName, valid := checkOrRespondError(w, req)
	if !valid {
		return nil
	}

	// 2. get sts bucket
	bucket, err := _oss.BuildOssStsBucket(ak, sk, arn)
	if err != nil && util.RespondError(constants.InternalError, w, err) {
		return nil
	}

	// 3. do signature for temporary
	signedUrl, err := bucket.SignURL(objectName, oss.HTTPGet, 5*60)
	if err != nil && util.RespondError(constants.InternalError, w, err) {
		return nil
	}

	return json.NewEncoder(w).Encode(&model.R{
		ErrCode: constants.Success,
		Url:     signedUrl,
	})
}

func checkOrRespondError(w http.ResponseWriter, req *http.Request) (string, string, string, string, bool) {
	ak, sk, objectName, valid := _func.CheckOrRespondError(w, req)
	if !valid {
		return "", "", "", "", false
	}

	arn := os.Getenv("OSS_STS_ROLE_ARN")
	if arn == "" {
		util.RespondError(constants.InternalError, w, fmt.Errorf("arn is empty!"))
		return "", "", "", "", false
	}

	return ak, sk, arn, objectName, true
}
