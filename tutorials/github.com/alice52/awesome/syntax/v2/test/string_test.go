package test

import (
	"testing"

	_str "github.com/alice52/awesome/syntax/v2/string"
)

func TestTraverse(t *testing.T) {

	s := "hello 世界"
	_str.Traverse(s)
}
