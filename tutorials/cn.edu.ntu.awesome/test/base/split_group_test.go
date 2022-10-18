package base

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 组测试 + 子测试 + 并行测试 + assert
func TestSplitV4(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := Split(tt.input, tt.sep)
			assert.Equal(got, tt.want)
		})
	}
}

// 组测试 + 子测试 + 并行测试
//
// Deprecated: use TestSplitV4 instead
func TestSplitV3(t *testing.T) {
	t.Parallel() // 将 TLog 标记为能够与其他测试并行运行
	// 定义测试表格
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{"base case", "a:b:c", ":", []string{"a", "b", "c"}},
		{"wrong sep", "a:b:c", ",", []string{"a:b:c"}},
		{"more sep", "abcd", "bc", []string{"a", "d"}},
		{"leading sep", "沙河有沙又有河", "沙", []string{"", "河有", "又有河"}},
	}
	// 遍历测试用例
	for _, tt := range tests {
		tt := tt                            // 注意这里重新声明tt变量
		t.Run(tt.name, func(t *testing.T) { // 使用t.Run()执行子测试
			t.Parallel() // 将每个测试用例标记为能够彼此并行运行
			got := Split(tt.input, tt.sep)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expected:%#v, got:%#v", tt.want, got)
			}
		})
	}
}

// 组测试 + 子测试
//
// Deprecated: use TestSplitV4 instead
func TestSplitV2(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

// 组测试
//
// Deprecated: use TestSplitV4 instead
func TestSplitV1(t *testing.T) {
	// 定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	TestSplitV0(t)
	// 定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	// 遍历切片，逐一执行测试用例
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%#v, got:%#v", tc.want, got)
		}
	}
}

// 单元测试用例
//
// Deprecated: use TestSplitV4 instead
func TestSplitV0(t *testing.T) {
	defer setupTestCase(t)(t)
	got := Split("a:b:c", ":")
	expect := []string{"a", "b", "c"}
	if !reflect.DeepEqual(expect, got) {
		t.Errorf("expected:%v, got:%v", expect, got)
	}
}
