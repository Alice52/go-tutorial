package test

import (
	"fmt"
	"testing"

	_array "github.com/alice52/awesome/syntax/v2/array"
)

var a int = 100
var (
	v int = 100
)

func init() {
	fmt.Printf("v: %v\n", v)

	var slice []int
	if slice == nil {
		fmt.Println("empty slice will be init as nil")
	}
	slice = append(slice, 1)
	fmt.Println("empty slice will be init as nil")

	fmt.Printf("slice: %v\n type %T", slice, slice)

	var slice2 [3]int
	fmt.Printf("slice2: %v\n type %T", slice2, slice2)

	s5 := []int{2}
	fmt.Printf("slice2: %v; type %T\n", s5, s5)

	s6 := make([]int, 0, 0)
	s6 = append(s6, 2, 3)
	fmt.Printf("s6: %v\n", s6)

	s7 := *new([]int)
	s7 = append(s7, 2, 3)
	fmt.Printf("s7: %v\n", s7)

	a := fmt.Sprintln("200")
	fmt.Printf("a: %T\n", a)

	b := fmt.Sprintf("%d", 200)
	fmt.Printf("b: %T\n", b)

	c, _ := fmt.Print(100)
	fmt.Printf("c: %T\n", c)

}

func TestInit(t *testing.T) {
	_array.InitArray()
}

func TestArrayValues(t *testing.T) {

	_array.ArrayValues()
}

func TestType(t *testing.T) {

	_array.DistinctType()
}
