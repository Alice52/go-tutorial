package test

import (
	_func "cn.edu.ntu.awesome/v0/syntax/func"
	"testing"
)

func TestAdd(t *testing.T) {
	_func.ChangeString()
	t.Logf("test add succ")
}
