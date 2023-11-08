package goroutine

import "sync"

// sync.Mutex
var (
	x  int64
	wg sync.WaitGroup // 等待组
	m  sync.Mutex     // 互斥锁
)
