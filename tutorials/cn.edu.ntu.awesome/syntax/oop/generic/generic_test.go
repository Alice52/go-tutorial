package generic

import (
	"testing"
)

func TestGeneric(t *testing.T) {
	// 1. 声明Set
	// 这里只能是 Animal: 内容可以存放实现
	// 与 java 的 List<Person> p = new ArrayList<Student>(); 编译不过是一样的
	var set Set[Animal] = &HashSet[Animal]{}

	// 2. 向Set中添加Dog
	var dog *Dog = &Dog{
		&Nameable{
			Name: "dog",
		},
	}
	set.Put(dog)
	dog.Bite()
	dog.GetName()

	// 3. 向Set中添加Cat
	var cat *Cat = &Cat{
		&Nameable{
			Name: "cat",
		},
	}
	set.Put(cat)

	// 4. 获取所有 pet
	for _, v := range set.All() {
		v.GetName()
	}

	// 6. pet
	pet := &Pet{
		Dog: dog,
		Cat: cat,
	}
	pet.Dog.Name = "Cat1"
	pet.Dog.GetName()
	pet.Dog.Bite()
}
