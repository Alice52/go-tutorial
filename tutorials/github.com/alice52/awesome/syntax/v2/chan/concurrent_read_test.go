package _chan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestConcurrentRead(t *testing.T) {
	var wg sync.WaitGroup
	defer wg.Wait()
	ch := make(chan int, 500)

	// 1. 启动多个 goroutine 读取数据
	routine(&wg, 10, ch)

	// 2. 向 chan 内写数据
	for i := 0; i < 100; i++ {
		ch <- i
	}

	// 3. 关闭 channel
	close(ch)
}

func routine(wg *sync.WaitGroup, count int, ch chan int) {
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				data, ok := <-ch
				if !ok {
					fmt.Printf("Goroutine %d: Channel closed\n", id)
					return
				}

				time.Sleep(1 * time.Second)
				fmt.Printf("Goroutine %d: Received %d\n", id, data)
			}
		}(i)
	}
}
