package _defer

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDeferRecursive(t *testing.T) {
	x := 1
	y := 2
	defer Calc("AA", x, Calc("A", x, y))
	x = 10
	defer Calc("BB", x, Calc("B", x, y))
	y = 20
}

func TestDefer(t *testing.T) {

	v1 := F1()
	v2 := F2()
	v3 := F3()
	v4 := F4()

	fmt.Printf("v1: %v\n; v2: %v\n v3: %v\n v4: %v\n", v1, v2, v3, v4)
}

func TestStacks(t *testing.T) {
	StackStore() // 3 2 1
}

func TestReturn(t *testing.T) {
	// 0
	log.Printf("%v", AnonymousReturn())
	// 1
	log.Printf("%v", NamedReturn())
}

func TestLoopDefer(t *testing.T) {
	DeferInLoop()
}

func TestIoClose(t *testing.T) {
	resp, err := IoClose("http://49.235.91.10:9200/")
	if err != nil {
		log.Fatalln(err.Error())
	}
	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("response status: %d, response: %v", resp.StatusCode, data)
}

func TestOsExit(t *testing.T) {
	OsExit()
}
