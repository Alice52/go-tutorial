package slice

import (
	"fmt"
	"log"
)

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
	var s6 []int
	s6 = arr[1:4]
	fmt.Println(s6)

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
