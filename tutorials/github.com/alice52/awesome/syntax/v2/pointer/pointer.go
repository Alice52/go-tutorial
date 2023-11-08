package pointer

import "fmt"

func Concept() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
	fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
	fmt.Println(&b)                    // 0xc00000e018
	fmt.Println(*b)                    // 10
}

func CreateReferenceTypeError() {
	var a1 *int // nil
	*a1 = 100   // error  a = new(int)
	fmt.Println(*a1)

	var b1 map[string]int
	b1["a"] = 100 // panic: assignment to entry in nil map
	fmt.Println(b1)
}

func DefineThenAllocate() {
	a1 := new(int)
	fmt.Printf("%T\n", a1) // *int
	*a1 = 100
	fmt.Println(*a1)

	var b map[string]int
	b = make(map[string]int, 10)
	b["a"] = 100
	fmt.Println(b)
}
