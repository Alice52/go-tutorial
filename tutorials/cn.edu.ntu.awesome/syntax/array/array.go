package array

import (
	"fmt"
	"log"
)

func InitArray() {
	var arrq [5]int
	fmt.Printf("arrq: %v\n", arrq)

	var arr0 [5]int = [5]int{1, 2, 3}
	var arr1 = [5]int{1, 2, 3, 4}
	var arr2 = [...]int{1, 2}
	var str = [5]string{3: "hello", 4: "world"}
	log.Printf("arr0: %v, arr1: %v, arr2: %v, str: %v", arr0, arr1, arr2, str)

	a := [3]int{1, 2, 3}
	b := [...]int{1, 2}
	c := [5]int{2: 10, 4: 9200}
	log.Printf("a: %v, b: %v, c: %v\n", a, b, c)

	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	log.Println(d)

	var arr20 = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	// var arr20 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	log.Println(arr20)

	var nArr = new([5]int)
	fmt.Printf("nArr: %v\n", nArr)
}

func ArrayValues() {
	arrA := [3]int{1, 2, 3}
	var arrB = arrA
	fmt.Printf("pStr: %p\n", &arrA) // pStr: 0xc00000e400
	fmt.Println(arrA == arrB)       // true

	arrA[0] = 20
	fmt.Printf("pStr: %p\n", &arrA) // pStr: 0xc00000e400
	fmt.Printf("pStr: %p\n", &arrB) // pStr: 0xc00000e420
	fmt.Println(arrA)               // [20 2 3]
	fmt.Println(arrB)               // [1 2 3]
	fmt.Println(arrA == arrB)       // false

	var arrC = new([5]int)
	arrD := arrC
	arrC[2] = 100
	fmt.Println(arrC)         // &[0 0 100 0 0]
	fmt.Println(arrD)         // &[0 0 100 0 0]
	fmt.Println(&arrC)        // 0xc000006030
	fmt.Println(&arrD)        // 0xc000006038
	fmt.Println(&arrC[0])     // 0xc00001a180
	fmt.Println(&arrD[0])     // 0xc00001a180
	fmt.Println(arrC == arrD) // true
}
