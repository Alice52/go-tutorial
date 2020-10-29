package test

import (
	_func "cn.edu.ntu.awesome/v0/syntax/func"
	"log"
	"testing"
)

func init() {
	log.Print("init execute ...")
}

func TestAdd(t *testing.T) {
	_func.ChangeString()

	t.Logf("test add succ")
}
