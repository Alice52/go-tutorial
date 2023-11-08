package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestSend(t *testing.T) {
	ch := make(chan int)
	go Recv(ch) // 创建一个 goroutine 从通道接收值
	Send(ch)
}

func TestDirection(t *testing.T) {
	Direction()
}

func TestMisuse(t *testing.T) {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func(no int) {
			for task := range ch {
				fmt.Printf("%d rece: %v\n", no, task)
			}
			wg.Done()
		}(j)
	}
	wg.Wait()
}
