package monkey

import (
	"reflect"
	"strings"
	"testing"

	"bou.ke/monkey"
	"cn.edu.ntu.awesome/test/monkey/lib"
)

func TestMyFunc(t *testing.T) {
	// 对 varys.GetInfoByUID 进行打桩
	monkey.Patch(lib.GetInfoByUID, func(int64) (*lib.UserInfo, error) {
		return &lib.UserInfo{Name: "liwenzhou"}, nil
	})

	ret := MyFunc(123)
	if !strings.Contains(ret, "liwenzhou") {
		t.Fatal()
	}
}

func TestUserMethod(t *testing.T) {
	var u = &User{
		Name:     "q1mi",
		Birthday: "1990-12-20",
	}

	// 为对象方法打桩
	monkey.PatchInstanceMethod(reflect.TypeOf(u), "CalcAge", func(*User) int {
		return 18
	})

	ret := u.CalcAge()
	if ret != 18 {
		t.Fatal()
	}
}
