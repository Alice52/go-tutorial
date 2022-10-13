package oop

import (
	"fmt"
)

type Address struct {
	province string
	city     string
}

type Person struct {
	name string
	age  int8
	addr *Address
}

type Student struct {
	*Person // 匿名嵌套对象(实现继承)
	school  string
	grade   string
}

func NewPerson(name string, age int8, addr *Address) *Person {
	return &Person{
		name: name,
		age:  age,
		addr: addr,
	}
}

/*********************************************
 * 使用指针类型所谓接受者
 *    1. 需要修改接收者中的值
 *    2. 接收者是拷贝代价比较大的大对象
 *    3. 保证一致性: 如果有某个方法使用了指针接收者, 那么其他的方法也应该使用指针接收者
 **********************************************/
func (p *Person) Saying(word string) *Person {
	fmt.Printf("%s say %s", p.name, word)
	return p
}

func (s *Student) Learning() {
	fmt.Printf("%s is learning.", s.name)
}

// 匿名字段默认使用类型名作为字段名
func AsAnonymousFile() {
	// User 用户结构体
	type User struct {
		Name    string
		Gender  string
		Address // 匿名字段
	}

	var user2 User
	user2.Address.province = "山东"
	user2.province = "山东" // 匿名字段可以省略
}

func Usage() {
	// 1. 对象 & 嵌套
	addr := &Address{
		province: "jx",
		city:     "xz",
	}
	p1 := NewPerson("zack", 18, addr)
	p1.addr.city = "nj"
	p1.Saying("person")

	// 2. 继承
	stu := &Student{
		Person: p1,
		grade:  "2",
		school: "nt",
	}

	stu.Saying("student")
	stu.Learning()

	// 3. 匿名字段
	AsAnonymousFile()

	// 4. 重载
	OverrideUsage()
}
