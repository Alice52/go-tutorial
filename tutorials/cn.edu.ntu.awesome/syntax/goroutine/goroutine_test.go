package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestLeak(t *testing.T) {
	Leak()
}

func TestContext(t *testing.T) {
	Context()
	time.Sleep(time.Second)
}

func TestHello(t *testing.T) {
	wg.Add(1)
	go Hello()

	fmt.Println("你好")
	wg.Wait()
	// time.Sleep(2 * time.Second)
}
