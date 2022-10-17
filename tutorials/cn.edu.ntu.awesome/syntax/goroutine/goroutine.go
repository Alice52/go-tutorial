package goroutine

import (
	"fmt"
	"time"
)

func Leak() {

	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "job result" // 2. 此时没有读取, 导致此 goroutine 阻塞
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second): // 1. 命中这里, 后结束 select
		fmt.Println("wait")
		return
	}
}

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
