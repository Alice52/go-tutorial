package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Add 对全局变量x执行 5000 次加 1 操作
func Add() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}

type Counter interface {
	Inc()
	Load() int64
}

var _ Counter = &CommonCounter{}
var _ Counter = &MutexCounter{}
var _ Counter = &AtomicCounter{}

// 普通版
type CommonCounter struct {
	counter int64
}

func (c *CommonCounter) Inc() {
	c.counter++
}
func (c *CommonCounter) Load() int64 {
	return c.counter
}

// 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func Concurrent(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}
