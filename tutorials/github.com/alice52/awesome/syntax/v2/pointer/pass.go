// Package pointer(参数等传递都是按值传递)
//  1. 传实例: 真正传递的是实例副本的地址 + 不可以修改实例的字段
//  2. 传指针: 指针的副本的地址(指向该实例) + 只可以修改实例的字段
//  3. 传指针的指针: 指针的副本的地址(指向前一个指针) + 可以对实例重新赋值(`*p=nil`)
package pointer

import "fmt"

type person struct {
	name string
	age  int
}

// addressInstance 此时的p是实例的副本, 无法修改实例的字段(三个示例对象的地址都不一样)
func addressInstance(p person) {
	fmt.Println(p)         // {alice 30}
	fmt.Println(&p)        // &{alice 30}
	fmt.Printf("%p\n", &p) // 0xc000008120

	var r person = p
	fmt.Println(&r)        // &{alice 30}
	fmt.Printf("%p\n", &r) // 0xc000008168
}

func addressPointer(p *person) {
	fmt.Println(p) // &{alice 30}
	// 查看这个指针变量(存储)指向的地址
	fmt.Println(&p)        // 0xc00000a058
	fmt.Printf("%p\n", &p) // 0xc00000a058
	fmt.Println(*p)        // {alice 30}

	var r *person = p
	fmt.Println(&r) // 0xc00000a058
}

// passInstance 此时传递的是p对象的副本(本质也是地址), 无法修改p对象的字段
func passInstance(p person) {
	p.name = "bob" // 只在此作用域内有效
}

// passPointer 此时传递的是指针的副本地址(也指向p对象的地址): 可以修改p对象的相关属性 + 不能将p对象重新赋值
func passPointer(p *person) {
	p.name = "bob" // 有效改动

	p = &person{
		name: p.name,
		age:  18,
	} // 无效改动
}

// passPointerPointer 此时传递的是指针的指针的副本地址(指向前一个指针的地址): 可以修改p对象的相关属性和将p对象重新赋值
func passPointerPointer(p **person) {
	(*p).name = "bob" // 有效改动

	*p = &person{
		name: (*p).name,
		age:  18,
	} // 有效改动
}
