package main

import (
	_func "github.com/alice52/proxy/sts/func"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {

	fc.StartHttp(_func.HandleHttpRequest)
}
