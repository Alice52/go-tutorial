package test

import (
	"fmt"
	"log"
	"testing"

	_func "cn.edu.ntu.awesome/v2/syntax/func"
)

func init() {
	log.Print("init execute ...")
}

func TestAdd(t *testing.T) {
	_func.ChangeString()

	t.Logf("test add succ")
}

func TestAsArg(t *testing.T) {
	v := _func.FasArg(1, 5, func(i1, i2 int) int { return i1 + i2 })
	fmt.Printf("v: %v\n", v)
}
