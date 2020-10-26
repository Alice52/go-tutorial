package _defer

import (
	"log"
	"net/http"
	"os"
)

/**
1. defer is stored in stack
2. function is same as using in .netcore
    + wraps defer function return will trigger defer execute
	+ wraps defer function execute to last line will trigger defer execute
    + this goroutine occurs exception will trigger defer execute
*/

// ***************************** 1 ***************************
// 3 2 1
func StackStore() {
	defer func() {
		log.Printf("1")
	}()

	defer func() {
		log.Printf("2")
	}()

	defer func() {
		log.Printf("3")
	}()
}

// ***************************** 2 ***************************
// 0
func AnonymousReturn() int {
	var result int
	defer func() {
		result++
		log.Println("defer execute...")
	}()

	return result
}

// 1
func NamedReturn() (result int) {
	defer func() {
		result++
		log.Println("defer execute...")
	}()

	return result
}

// ***************************** 3 ***************************
// warning: Possible resource leak, 'defer' is called in a 'for' loop
func DeferInLoop() {
	for i := 0; i < 100; i++ {
		f, _ := os.Open("../go.mod")
		defer f.Close()
	}
}

// ***************************** 4 ***************************
func IoClose(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	// do this first
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}

// ***************************** 5 ***************************
// os.Exit will not trigger defer execute
func OsExit() {
	defer func() {
		log.Println("defer execute...")
	}()
	os.Exit(0)
}
