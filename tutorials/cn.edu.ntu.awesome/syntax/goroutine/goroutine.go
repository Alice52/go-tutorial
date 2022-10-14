package goroutine

import (
	"fmt"
	"sync"
)

// sync.Mutex
var (
	x  int64
	wg sync.WaitGroup // 等待组
	m  sync.Mutex     // 互斥锁
)

func Context() {
	for i := 0; i < 5; i++ {
		go func() {
			// all respose is 5
			fmt.Printf("loopclosure value: %v\n", i)
		}()
	}

	for i := 0; i < 5; i++ {
		// 要执行的函数 + 上下文
		go func(i int) {
			fmt.Printf("go func value: %v\n", i)
		}(i)
	}
}

func Hello() {
	defer wg.Done()

	fmt.Println("Hello Goroutine")
}

// add 对全局变量x执行 5000 次加 1 操作
func add() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}
