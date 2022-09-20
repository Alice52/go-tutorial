package inherit

import "fmt"

type Skills []string

type person struct {
	name   string
	age    int
	weight int
}

type Student struct {
	person //继承
	Skills
	int
	spe string
}

func Inherit() {
	//方式一,全部指定
	xuxu := Student{person{"xuxu", 25, 68}, []string{"anatomy"}, 1, "boy"}
	//方式二,指哪打哪
	jane := Student{person: person{"Jane", 25, 100}, spe: "Biology"}

	fmt.Printf("His name is %s", jane.name)
	fmt.Printf("His name is %s", xuxu.name)
}
