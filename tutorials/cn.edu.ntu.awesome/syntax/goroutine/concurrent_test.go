package goroutine

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	wg.Add(2)

	go Add()
	go Add()

	wg.Wait()
	fmt.Println(x)
}

func TestCocurrent(t *testing.T) {
	c1 := CommonCounter{} // 非并发安全
	Concurrent(&c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	Concurrent(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	Concurrent(&c3)
}
