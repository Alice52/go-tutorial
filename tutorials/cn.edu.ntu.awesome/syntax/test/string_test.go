package test

import (
	"testing"

	_str "cn.edu.ntu.awesome/v0/syntax/string"
)

func TestTraverse(t *testing.T) {

	s := "hello 世界"
	_str.Traverse(s)
}
