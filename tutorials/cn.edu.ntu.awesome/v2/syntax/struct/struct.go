package _struct

import (
	"encoding/json"
	"fmt"
	"log"

	"cn.edu.ntu.awesome/v2/syntax/common"
)

// will init as fileds zero value.
var person common.Person

func StructTag() {

	type Student struct {
		ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
		Gender string //json序列化是默认使用字段名作为key
		name   string //私有不能被json包访问
	}
	s1 := &Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, _ := json.Marshal(s1)
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}

func NewPerson(name, city string, age int8) *common.Person {

	return &common.Person{
		Name: name,
		City: city,
		Age:  age,
	}
}

func StructPtr() {
	m := make(map[string]*common.Student)
	stus := []common.Student{
		{Name: "小王子", Age: 18},
		{Name: "娜扎", Age: 23},
		{Name: "大王八", Age: 9000},
	}

	// 这里 stu 值是不一样的, 但是地址是一直复用的
	for _, stu := range stus {
		fmt.Printf("stu: %p\n", &stu)
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
}

func Initialized() {
	// 1. declare and use
	// var person Person
	person.Age = 15
	person.Name = "zack"
	log.Printf("person: %#v\n", person)

	// 2. pointer
	p := new(common.Person)
	p.Name = "kayla"
	p.Age = 15
	log.Printf("type: %T, addr: %p,  person: %#v\n", p, p, *p)

	// 3. initialize
	p5 := common.Person{
		Name: "prof.cn",
		City: "BeiJing",
		Age:  18,
	}
	log.Printf("type: %T, addr: %p,  person: %#v\n", p5, &p5, p5)
	pp := &p5
	log.Printf("type: %T, addr: %p,  person: %#v\n", pp, pp, *pp)

	// 4. initialize and pointer
	p6 := &common.Person{
		Name: "prof.cn",
		City: "BeiJing",
		Age:  18,
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
