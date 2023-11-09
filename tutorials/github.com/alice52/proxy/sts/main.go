package main

import (
	_func "github.com/alice52/proxy/sts/func"
	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

// 使用 STS 机制对Ak进行保护
// 调用 AssumeRole 接口有限流要求: 因此需要保存 SecuriyToken(有状态服务)
// 因此不适合作为函数(可以缓存在 redis 内实现)
func main() {

	fc.StartHttp(_func.HandleHttpRequest)
}
