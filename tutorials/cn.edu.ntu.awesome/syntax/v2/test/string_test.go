package test

import (
	"testing"

	_str "cn.edu.ntu.awesome/syntax/v2/string"
)

func TestTraverse(t *testing.T) {

	s := "hello 世界"
	_str.Traverse(s)
}
