package goroutine

import (
	"sync"
)

// sync.Mutex

var (
	x  int64
	wg sync.WaitGroup // 等待组
	m  sync.Mutex     // 互斥锁
)

// add 对全局变量x执行 5000 次加1操作
func add() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}
