package common

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Person struct {
	Name string
	City string
	Age  int8
}

type (
	Rect struct {
		width  float64
		height float64
	}
)

func CreateObj() {

	var r1 *Rect
	r1 = new(Rect) // make create slice/map/channel
	r1.height = 1
	fmt.Println(r1) // &{0 1}

	r3 := &Rect{
		width: 1,
	}
	fmt.Println(r3) // &{1 0}

	var r2 Rect = Rect{width: 1} // r2 := Rect{width: 1}
	fmt.Println(r2)              // {1 0}
}

// *取变量的值
// &取变量的地址
func Operator() {
	rect := &Rect{100, 100}

	// 取到Rect类型对象的地址
	fmt.Println(rect) // &{100 100}
	// 查看这个指针变量的地址
	fmt.Println(&rect) // 0xc000098028
	fmt.Println(*rect) // {100 100}

	var r *Rect = rect
	// *表示变量指针类型
	fmt.Println(r)  // &{100 100}
	fmt.Println(&r) // 0xc000098028

	fmt.Println(*r) // {100 100}

}

func (r *Rect) area() float64 {
	return r.width * r.height
}
