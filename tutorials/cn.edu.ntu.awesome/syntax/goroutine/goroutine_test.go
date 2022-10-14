package goroutine

import (
	"fmt"
	"testing"
	"time"
)

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

func Test_add(t *testing.T) {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
