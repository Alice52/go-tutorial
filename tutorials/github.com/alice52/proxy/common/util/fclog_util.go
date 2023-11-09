package util

import (
	"context"
	"fmt"
	"github.com/aliyun/fc-runtime-go-sdk/fccontext"
	"net/http"
)

func LogFcContext(ctx context.Context, req *http.Request) {

	lc, _ := fccontext.FromContext(ctx)
	fmt.Printf("context: %#v\n", lc)
	fmt.Printf("req: %#v\n", req)
}
