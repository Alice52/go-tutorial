package goroutine

import (
	"fmt"
	"strconv"
	"sync"
)

var mp = make(map[string]int)

// 并发安全的map
var mp_safe = sync.Map{}

func SafeMap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			mp_safe.Store(key, n)         // 存储key-value
			value, _ := mp_safe.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func UnsafeMap() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			mp[key] = n
			fmt.Printf("k=:%v,v:=%v\n", key, mp[key])
			wg.Done()
		}(i)
	}
	wg.Wait()
}
