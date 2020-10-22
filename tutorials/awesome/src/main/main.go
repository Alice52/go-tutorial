package main

import "fmt"

const pi = 3.1415926
const (
	Sunday = iota
	Monday
	Tuesday
	_
	Thursday
	Friday
	Saturday
)

func main() {

	// array is pass by value, change value will do not allocate new space.
	// slice is pass by reference, and slice is left-closed and right-opened interval.
	arr := [3]int{1, 2, 3}
	fmt.Printf("%T\n", arr) // [3]int

	for _, v := range arr {
		fmt.Println(v)
	}

	sli := arr[1:3]         // [2 3]
	fmt.Printf("%T\n", sli) // []int

	// change sli, so arr will change too
	sli[1] = 5
	fmt.Println(arr) 				// [1 2 5]
	fmt.Println(sli) // [2 5]

	// change arr will not allocate new space, so it will have impact on slice
	arr[1] = 10
	fmt.Println(arr) // [1 10 5]
	fmt.Println(sli) // [10 5]

	arrA := [3]int{1, 2, 3}
	var arrB = arrA
	fmt.Printf("pStr: %p\n", &arrA) // pStr: 0xc00000e400

	arrA[0] = 20
	fmt.Printf("pStr: %p\n", &arrA) // pStr: 0xc00000e400
	fmt.Printf("pStr: %p\n", &arrB) // pStr: 0xc00000e420
	fmt.Println(arrA)               // [1 2 3]
	fmt.Println(arrB)               // [1 2 3]

	var arrC = new([5]int)
	arrD := arrC
	arrC[2] = 100
	fmt.Println(arrC)  // &[0 0 100 0 0]
	fmt.Println(arrD)  // &[0 0 100 0 0]
	fmt.Println(&arrC) // 0xc000006030
	fmt.Println(&arrD) // 0xc000006038

	// Slice
	sliA := []int{0, 1, 2, 3, 4, 5}
	sliB := sliA[2:4]
	fmt.Println(sliB)      // [2 3]
	fmt.Println(len(sliB)) // 2
	fmt.Println(cap(sliB)) // 4
	fmt.Println(sliB[1])   // 3
	//fmt.Println(sliB[2])                                          // error
	sliC := sliB[:4]
	fmt.Println(sliC) // [2 3 4 5]

	var s1 []int
	fmt.Printf("len: %d; cap: %d.\n", len(s1), cap(s1))

	// point
	i := 42
	p := &i         // Set 'p' points to 'i'
	fmt.Println(*p) // 42
	*p = 21         // Set 'i' through the pointer 'p'
	fmt.Println(*p) // 21
	fmt.Println(i)  // 21

	*p = *p / 3
	fmt.Println(*p) // 7
	fmt.Println(i)  // 7

	// const: iota will add 1 per const row
	fmt.Println(Friday) // 5

	changeString()
	switchDemo1() // 3 4

	// for range

	for {
		// break goto return panic to exit
		fmt.Println(1)
		break
	}
}

func changeString() {
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("1") // default has break
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		// used in switch, will execute following express
		// no matter what following condition
		fallthrough
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("6")
	}
}
