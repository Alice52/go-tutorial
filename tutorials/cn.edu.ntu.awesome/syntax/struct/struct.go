package _struct

import (
	"log"
)

// will init as fileds zero value.
var person Person

type Person struct {
	name string
	city string
	age  int8
}

func Initialized() {
	// 1. declare and use
	// var person Person
	person.age = 15
	person.name = "zack"
	log.Printf("person: %#v\n", person)

	// 2. pointer
	p := new(Person)
	p.name = "kayla"
	p.age = 15
	log.Printf("type: %T, addr: %p,  person: %#v\n", p, p, *p)

	// 3. initialize
	p5 := Person{
		name: "prof.cn",
		city: "BeiJing",
		age:  18,
	}
	log.Printf("type: %T, addr: %p,  person: %#v\n", p5, &p5, p5)
	pp := &p5
	log.Printf("type: %T, addr: %p,  person: %#v\n", pp, pp, *pp)

	// 4. initialize and pointer
	p6 := &Person{
		"prof.cn",
		"BeiJing",
		18,
	}
	log.Printf("type: %T, addr: %p,  person: %#v\n", p6, p6, *p6)
}

func AnonymousStruct() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "zack"
	user.Age = 15
	log.Printf("user: %#v\n", user)
}
