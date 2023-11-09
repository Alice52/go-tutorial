package _func

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/alice52/proxy/common/component"
	"github.com/alice52/proxy/common/log"
	"github.com/alice52/proxy/common/resp"
	_func "github.com/alice52/proxy/oss/func"
	"github.com/alice52/proxy/oss/model"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"net/http"
	"os"
)

// HandleHttpRequest check env, params and signature temporary image url for  by oss bucket
func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {

	log.LogFcContext(ctx, req)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// 1. check oss env and request params, and respond error if invalid
	ak, sk, arn, objectName, valid := checkOrRespondError(w, req)
	if !valid {
		return nil
	}
	fmt.Println(ak, sk, arn, objectName, valid)

	// 2. get sts bucket
	bucket, err := component.BuildOssStsBucket(ak, sk, arn)
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

func checkOrRespondError(w http.ResponseWriter, req *http.Request) (string, string, string, string, bool) {
	ak, sk, objectName, valid := _func.CheckOrRespondError(w, req)
	if !valid {
		return "", "", "", "", false
	}

	arn := os.Getenv("OSS_STS_ROLE_ARN")
	if arn == "" {
		resp.RespondError(w, fmt.Errorf("arn is empty!"))
		return "", "", "", "", false
	}

	return ak, sk, arn, objectName, true
}
