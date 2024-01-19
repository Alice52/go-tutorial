package _struct

import (
	"fmt"
	"testing"
)

func TestTag(t *testing.T) {
	StructTag()
}

func TestConstructor(t *testing.T) {
	p := NewPerson("zack", "xz", 18)
	fmt.Printf("p: %v, type: %T\n", p, p)
}

func TestStruct(t *testing.T) {
	StructPtr()
}

func TestAlias(t *testing.T) {
	AliasInfo()
}

func TestInitial(t *testing.T) {
	Initialized()
}

func TestAnonymousStruct(t *testing.T) {
	AnonymousStruct()
}
