package convey

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// 组测试 + 子测试 + 并行测试 + assert(convey)
func TestSplit(t *testing.T) {
	// t.Parallel()

	tests := []struct {
		name   string
		input  string
		sep    string
		expect []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	for _, tt := range tests {
		tt := tt // important

		convey.Convey(tt.name, t, func() {
			got := Split(tt.input, tt.sep)
			convey.So(got, convey.ShouldResemble, tt.expect)
		})
	}
}
