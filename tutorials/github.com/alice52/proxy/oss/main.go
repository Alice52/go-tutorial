package main

import (
	_func "github.com/alice52/proxy/oss/func"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	fc.StartHttp(_func.HandleHttpRequest)
}
