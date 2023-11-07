package main

import (
	"github.com/aliyun/fc-runtime-go-sdk/fc"
	_func "io.github.alice52.proxy/oss/func"
)

func main() {

	fc.StartHttp(_func.HandleHttpRequest)
}
