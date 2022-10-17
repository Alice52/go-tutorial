package goroutine

import (
	"fmt"
	"testing"
)

func TestConsumer(t *testing.T) {
	ch := Producer()
	res := Consumer(ch)
	fmt.Println(res) // 25
}
