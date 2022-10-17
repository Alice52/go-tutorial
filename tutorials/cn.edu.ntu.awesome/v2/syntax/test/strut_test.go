package test

import (
	"fmt"
	"testing"

	_struct "cn.edu.ntu.awesome/v2/syntax/struct"
)

func TestTag(t *testing.T) {
	_struct.StructTag()
}

func TestConstructor(t *testing.T) {
	p := _struct.NewPerson("zack", "xz", 18)
	fmt.Printf("p: %v, type: %T\n", p, p)
}

func TestStruct(t *testing.T) {
	_struct.StructPtr()
}

func TestAlias(t *testing.T) {
	_struct.AliasInfo()
}

func TestInitial(t *testing.T) {
	_struct.Initialized()
}

func TestAnonymousStruct(t *testing.T) {
	_struct.AnonymousStruct()
}
