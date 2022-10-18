package goroutine

import "fmt"

// Producer 返回一个通道, 并向其中发送数据, 之后关闭通道
// func Producer() chan int {
func Producer() <-chan int { // 定义返回接收通道
	ch := make(chan int, 2)
	// 创建一个新的 goroutine 执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道(否则会直接报错{deadlock})
	}()

	return ch
}

// Consumer 从通道中接收数据进行计算
// func Consumer(ch chan int) int {
func Consumer(ch <-chan int) int { // 定义一个接受向的通道
	sum := 0
	for v := range ch {
		fmt.Printf("consumer rec value: %v\n", v)
		sum += v
	}

	return sum
}
