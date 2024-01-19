package conc

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"sync"
	"sync/atomic"
	"time"
)

func syncGroupDemo1() {
	var count atomic.Int64
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1) // 每启动一个 goroutine，计数器加一
		go func() {
			defer wg.Done()
			count.Add(1)
			if count.Load() == 5 {
				time.Sleep(10 * time.Second)
			}
		}()
	}

	wg.Wait()

	fmt.Println(count.Load())
}

// waitGroupDemo1 开启10个goroutine并发执行 count.Add(1)
// 1. Go方法能够生成goroutine
// 2. WaitGroup确保在生成的goroutine都退出后再继续执行
// 3. conc.WaitGroup 比 sync.WaitGroup 多将子goroutine中的panic会被传递给Wait方法的调用方(无需自己recovery goroutine)
func waitGroupDemo1() {
	var count atomic.Int64

	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			count.Add(1)
		})
	}
	// 等待10个goroutine都执行完
	wg.Wait()

	fmt.Println(count.Load())
}

// syncGroupDemo2 自动 recover示例
func syncGroupDemo2() {
	var count atomic.Int64
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1) // 每启动一个 goroutine，计数器加一
		go func() {
			// 在里面使用I是大忌讳
			defer wg.Done()
			count.Add(1)
			if count.Load() == 5 {
				panic("bad thing")
			}
		}()
	}

	wg.Wait()

	fmt.Println(count.Load())
}

// waitGroupDemo1 自动 recover示例
func waitGroupDemo2() {
	var count atomic.Int64

	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			// 在里面使用I是大忌讳
			if count.Load() == 7 {
				// panic("bad thing")
			}
			count.Add(1)
		})
	}

	r := wg.WaitAndRecover()
	if r != nil {
		fmt.Println(r)
	}
	fmt.Println(count.Load())
}
