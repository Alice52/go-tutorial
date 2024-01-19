package conc

import (
	"fmt"
	"github.com/sourcegraph/conc/iter"
)

func iterDemo() {
	input := []int{1, 2, 3, 4}

	iter.ForEach(input, func(v *int) {
		if *v%2 != 0 {
			*v = -1
		}
	})
	fmt.Println(input)
}

// iteratorDemo 创建一个最大goroutine个数为输入元素一半的迭代器
func iteratorDemo() {
	input := []int{1, 2, 3, 4}
	iterator := iter.Iterator[int]{
		MaxGoroutines: len(input) / 2,
	}

	iterator.ForEach(input, func(v *int) {
		if *v%2 != 0 {
			*v = -1
		}
	})

	fmt.Println(input)
}

// mapperDemo 创建一个最大goroutine个数为输入元素一半的映射器
func mapperDemo() {
	input := []int{1, 2, 3, 4}
	mapper := iter.Mapper[int, bool]{
		MaxGoroutines: len(input) / 2,
	}

	results := mapper.Map(input, func(v *int) bool { return *v%2 == 0 })
	fmt.Println(results)
}
