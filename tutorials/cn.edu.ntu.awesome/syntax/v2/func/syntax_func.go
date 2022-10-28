package _func

import "fmt"

// 1. as type
type Op func(int, int) int

// 2. as arg
func calc(x, y int, op Op) Op {

	// 3. as var
	opFunc := op
	opFunc(x, y)

	return func(a, b int) int {
		return a + b
	}
}

func main() {
	// as arg
	add := func(a, b int) int {
		return a + b
	}
	// excute func
	add(1, 2)

	// pass as arg
	calc(10, 20, add)
}

func ChangeString() {
	defer println("as")
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

func FasAnonymous() {
	// 将匿名函数保存到变量
	add := func(x, y int) int {
		fmt.Println(x + y)
		return x + y
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
