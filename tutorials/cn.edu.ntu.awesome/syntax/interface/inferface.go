package _interface

import "fmt"

// 定义接口
type Flyable interface {
	Fly()
}

type Bird struct {
	Name string
}

// impl check by compiler
var _ Flyable = &Bird{}

func (b *Bird) Fly() {
	fmt.Printf("%v is flying\n", b.Name)
}

func AssertType() {
	var fly Flyable = &Bird{"zack"}
	v, ok := fly.(*Bird)
	if ok {
		fmt.Println("fly 类型是 Bird")
		v.Name = "富贵"
	} else {
		fmt.Println("fly 类型不是 Bird")
	}
}

func JustifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}
}
