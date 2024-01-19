package main

import (
	"log"
	"testing"

	_func "github.com/alice52/awesome/syntax/v2/func"
)

func init() {
	log.Print("init execute ...")
}

func TestAdd(t *testing.T) {
	_func.ChangeString()

	t.Logf("test add succ")
}
