package conc

import (
	"context"
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/pool"
	"strconv"
)

func poolDemo() {
	// 创建一个最大数量为3的goroutine池
	p := pool.New().WithMaxGoroutines(3)
	// 使用p.Go()提交5个任务
	for i := 0; i < 5; i++ {
		i := i
		p.Go(func() {
			if i == 2 {
				// panic("i")
			}
			fmt.Println("Q1mi")
		})
	}
	p.Wait() // 此处会直接re-panic
}

// poolWithContextDemoCancelOnError 支持context的池(goroutine中出错时取消context)
func poolWithContextDemoCancelOnError() {
	p := pool.New().
		WithMaxGoroutines(4).
		WithContext(context.Background()).
		WithCancelOnError() // 出错时取消所有goroutine

	for i := 0; i < 20; i++ {
		i := i
		p.Go(func(ctx context.Context) error {
			fmt.Printf("start i: %d\n", i)
			if i == 2 {
				return errors.New("cancel all other tasks")
			}
			<-ctx.Done() // 确定一个上下文是否已经被取消
			fmt.Printf("end i: %d\n", i)
			return nil
		})
	}
	err := p.Wait()
	fmt.Println(err)
}

func poolWithError() {
	p := pool.New().WithErrors()
	for i := 0; i < 3; i++ {
		i := i
		p.Go(func() error {
			if i == 2 || i == 1 {
				return errors.New("oh no!" + strconv.Itoa(i))
			}
			return nil
		})
	}
	err := p.Wait() // 内部是多个err(不是)
	fmt.Println(err)
}

// poolWithResult 执行返回结果的任务池
func poolWithResult() {
	p := pool.NewWithResults[int]()
	for i := 0; i < 10; i++ {
		i := i
		p.Go(func() int {
			if i == 2 {
				// panic("panic")
			}
			return i * 2
		})
	}
	res := p.Wait()
	fmt.Println(res)
}
