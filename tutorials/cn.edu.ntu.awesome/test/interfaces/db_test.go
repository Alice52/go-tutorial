package interfaces

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"cn.edu.ntu.awesome/test/interfaces/mocks"
	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	// 创建 gomock 控制器
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 这里mocks是我们生成代码时指定的package名称
	m := mocks.NewMockDB(ctrl)
	// 打桩: 当传入Get函数的参数为 liwenzhou.com 时返回1和nil
	m.EXPECT().
		Get(gomock.Eq("liwenzhou.com")). // 参数
		Return(1, nil).                  // 返回值
		Times(1)                         // 调用次数

	// 调用GetFromDB函数时传入上面的mock对象m
	if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
		t.Fatal()
	}

	// 再次调用上方mock的Get方法时不满足调用次数为1的期望
	// if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
	// 	t.Fatal()
	// }

	// 指定顺序
	gomock.InOrder(
		m.EXPECT().Get("1"),
		m.EXPECT().Get("2"),
		m.EXPECT().Get("3"),
	)

	// 按顺序调用
	GetFromDB(m, "1")
	GetFromDB(m, "2")
	GetFromDB(m, "3")
}

func TestGetFromDB2(t *testing.T) {
	var db DB
	db = &MySQL{}
	fmt.Printf("reflect.TypeOf(db): %v\n", reflect.TypeOf(db).Kind())
	monkey.PatchInstanceMethod(reflect.TypeOf(db), "Get", func(db *MySQL, key string) (int, error) {
		return 100, nil
	})

	if v := GetFromDB(db, "liwenzhou.com"); v != 1100 {
		t.Fatalf("expect 100, got %v\n", v)
	}
}
