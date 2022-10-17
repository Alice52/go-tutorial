package _struct

import "fmt"

// new type, value is int
type NewInt int

// alias
type MyInt = int

func AliasInfo() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) // type of a:_struct.NewInt
	fmt.Printf("type of b:%T\n", b) // type of b:int
}
