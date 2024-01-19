package conc

import (
	"fmt"
	"github.com/sourcegraph/conc/stream"
	"time"
)

// streamDemo 并发的流式任务示例
func streamDemo() {
	times := []int{20, 52, 16, 45, 4, 80}

	s := stream.New()
	for _, millis := range times {
		dur := time.Duration(millis) * time.Millisecond
		s.Go(func() stream.Callback {
			time.Sleep(dur)
			if millis == 80 {
				// panic("panic")
			}
			// 虽然上一行通过sleep增加了时间, 但最终结果仍按任务提交（s.Go）的顺序打印
			return func() { fmt.Println(dur) }
		})
	}
	s.Wait() // re-panic if necessary
}
