package monkey

import (
	"fmt"
	"github.com/alice52/awesome/test/monkey/lib"
	"time"
)

func MyFunc(uid int64) string {
	u, err := lib.GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return fmt.Sprintf("hello %s\n", u.Name)
}

type User struct {
	Name     string
	Birthday string
}

// CalcAge 计算用户年龄
func (u *User) CalcAge() int {
	t, err := time.Parse("2006-01-02", u.Birthday)
	if err != nil {
		return -1
	}
	return int(time.Now().Sub(t).Hours()/24.0) / 365
}
