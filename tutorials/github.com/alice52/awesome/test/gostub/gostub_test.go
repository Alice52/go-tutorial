package gostub

import (
	"testing"

	"github.com/prashantv/gostub"
)

func TestConfig(t *testing.T) {
	// 打桩全局变量
	stubs := gostub.Stub(&configFile, "./test.toml")
	defer stubs.Reset()

	// 下面是测试的代码
	data, err := GetConfig()
	if err != nil {
		t.Fatal()
	}
	// 返回的data的内容就是上面/tmp/test.config文件的内容
	t.Logf("data:%s\n", data)
}

func TestShowNumber(t *testing.T) {
	stubs := gostub.Stub(&maxNum, 20)
	defer stubs.Reset()
	// 下面是一些测试的代码
	res := ShowNumber()
	if res != 20 {
		t.Fatal()
	}
}
