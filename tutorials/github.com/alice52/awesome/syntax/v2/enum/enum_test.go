package enum

import (
	"testing"
)

func TestIntEnum(t *testing.T) {
	Value(10) // 编译可以过(但是10不是枚举值)
}
