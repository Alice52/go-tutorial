package _defer

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

///  version2

func Calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

// 5
func F1() int {
	x := 5 // 1. x=5
	defer func() {
		x++ // 3. x=6, 但是返回值是5
	}()
	return x // 2. x=5, 此时就确定了
}

// 6
func F2() (x int) { // 1. x 最初是0
	defer func() {
		x++ // 3. x=6, 返回的是 x 变量
	}()
	return 5 // 2. x=5
}

// 5 与 F1() 一样
func F3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func F4() (x int) { // 1. x=0
	defer func(x int) {
		x++ // 4. x=1
	}(x) // 2. x=0 作为参数
	return 5 // 3. x=5
}

/// version1

// ***************************** 1 ***************************
// 3 2 1
func StackStore() {
	defer func() {
		log.Printf("1")
	}()

	defer func() {
		log.Printf("2")
	}()

	defer func() {
		log.Printf("3")
	}()
}

// ***************************** 2 ***************************
// 0
func AnonymousReturn() int {
	var result int
	defer func() {
		result++
		log.Println("defer execute...")
	}()

	return result
}

// 1
func NamedReturn() (result int) {
	defer func() {
		result++
		log.Println("defer execute...")
	}()

	return result
}

// ***************************** 3 ***************************
// warning: Possible resource leak, 'defer' is called in a 'for' loop
func DeferInLoop() {
	for i := 0; i < 100; i++ {
		f, _ := os.Open("../go.mod")
		defer f.Close()
	}
}

// ***************************** 4 ***************************
func IoClose(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	// do this first
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}

// ***************************** 5 ***************************
// os.Exit will not trigger defer execute
func OsExit() {
	defer func() {
		log.Println("defer execute...")
	}()
	os.Exit(0)
}
