package goroutine

import "fmt"

// i = 1 时, x := <-ch 不满足, 而执行 ch <- i
// i = 2 时, ch <- i 不满足(通道已满), 而执行 x := <-ch(接受数据)
func Odd() {
	var ch chan int = make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case data := <-ch:
			fmt.Printf("data: %v\n", data)
		case ch <- i:
		default:
		}
	}
}
