package slice

import (
	"fmt"
	"log"
)

var (
	s [10]int
)

func init() {
	s[0] = 1
}

func CreateSlice() []int {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int = arr[1:4]
	fmt.Println(s6)

	// init
	s7 := []int{1, 2, 3}
	s8 := s7
	s8[2] = 100
	fmt.Printf("s7: %v, type: %T\n", s7, s7)
	fmt.Printf("s8: %v, type: %T\n", s8, s8)

	var a1 = *new(int)
	a1 = 100
	fmt.Println(a1)

	var a2 = new(int)
	fmt.Println(a2)

	var a3 = new([]int) // nil
	fmt.Println(a3, append(*a3, 1))
	var a31 = *new([]int)
	a31 = append(a31, 1)
	fmt.Println(a31)

	var a4 *int
	// *a4 = 100 // panic
	fmt.Println(a4)

	var a5 = *new([]int)
	if a5 == nil {
		fmt.Println("a5 is nil")
	}
	fmt.Println(a5)

	return s6
}

func InitSlices() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[0:1]
	var slice1 []int = arr[:5]
	var slice2 []int = arr[1:]
	var slice3 []int = arr[:]
	var slice4 = arr[:len(arr)-1] //去掉切片的最后一个元素

	log.Println(slice0, slice1, slice2, slice3, slice4)

	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	slice5 := arr2[0:2]
	slice6 := arr2[:5]
	slice7 := arr2[4:]
	slice8 := arr2[:]
	slice9 := arr2[:len(arr)-1] //去掉切片的最后一个元素

	log.Println(slice5, slice6, slice7, slice8, slice9)
}
