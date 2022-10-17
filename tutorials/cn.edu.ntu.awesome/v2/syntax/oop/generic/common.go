package generic

import "fmt"

type Animal interface {
	GetName()
}

type Nameable struct {
	Name string
}

func (t *Nameable) GetName() {
	fmt.Printf("name: %v\n", t.Name)
}

type Dog struct {
	*Nameable
}

// impl check for compiler
var _ Animal = &Dog{}

func (t *Dog) Bite() {
	fmt.Printf("%v bite\n", t.Name)
}

type Cat struct {
	*Nameable
}

// impl check for compiler
var _ Animal = &Cat{}

func (t *Dog) Meow() {
	fmt.Printf("%v meow\n", t.Name)
}

type Pet struct {
	*Dog
	*Cat
}
