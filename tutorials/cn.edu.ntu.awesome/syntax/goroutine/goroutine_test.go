package goroutine

import (
	"fmt"
	"testing"
)

func Test_add(t *testing.T) {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
