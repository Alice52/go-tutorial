package pointer

import (
	"fmt"
	"testing"
)

var p = person{
	name: "alice",
	age:  30,
}

func Test_addressInstance(t *testing.T) {

	fmt.Printf("%p\n", &p) // 0xc000008108
	addressInstance(p)
}

func Test_addressPointer(t *testing.T) {

	addressPointer(&p)
}

func Test_passInstance(t *testing.T) {
	passInstance(p)
	t.Log(p)
}

func Test_passPointer(t *testing.T) {
	passPointer(&p)
	t.Log(p)
}

func Test_passPointerPointer(t *testing.T) {
	ptr := &p
	passPointerPointer(&ptr)
	t.Log(p)   // {bob 30}
	t.Log(&p)  // &{bob 30}
	t.Log(ptr) // &{bob 18}
}
