package test

import (
	_struct "cn.edu.ntu.awesome/v0/syntax/struct"
	"fmt"
	"log"
	"testing"
)

func TestAlias(t *testing.T) {
	_struct.AliasInfo()
}

func TestInitial(t *testing.T) {
	_struct.Initialized()
}

func TestAnonymousStruct(t *testing.T) {
	_struct.AnonymousStruct()
}

func TestStructUsage(t *testing.T) {
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	students := []*student{
		{name: "prof.cn", age: 18},
		{name: "testing", age: 23},
		{name: "blog", age: 28},
	}
	for _, stu := range students {
		fmt.Printf("student value: %v, addr: %p\n", stu, &stu)
		m[stu.name] = stu
	}
	a := 1

	b := &a
	c := &b
	fmt.Printf("student value: %p, addr: %d, %p\n", &b, **c, &a)

	for i := 0; i < 3; i++ {
		m[students[i].name] = students[i]
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	s1 := [...]int{
		2, 2, 3,
	}
	test1(&s1)
	log.Println(s1)

	s2 := []*int{
		&a,
	}
	test2(s2)
	log.Println(*s2[0])
}

func test1(a *[3]int) {
	a[0] = 1
}

func test2(a []*int) {
	b := 2
	a[0] = &b
}
