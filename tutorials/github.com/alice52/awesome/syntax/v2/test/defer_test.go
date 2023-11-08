package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	_defer "github.com/alice52/awesome/syntax/v2/defer"
)

func TestDeferRecursive(t *testing.T) {
	x := 1
	y := 2
	defer _defer.Calc("AA", x, _defer.Calc("A", x, y))
	x = 10
	defer _defer.Calc("BB", x, _defer.Calc("B", x, y))
	y = 20
}

func TestDefer(t *testing.T) {

	v1 := _defer.F1()
	v2 := _defer.F2()
	v3 := _defer.F3()
	v4 := _defer.F4()

	fmt.Printf("v1: %v\n; v2: %v\n v3: %v\n v4: %v\n", v1, v2, v3, v4)
}

func TestStacks(t *testing.T) {
	_defer.StackStore() // 3 2 1
}

func TestReturn(t *testing.T) {
	// 0
	log.Printf("%v", _defer.AnonymousReturn())
	// 1
	log.Printf("%v", _defer.NamedReturn())
}

func TestLoopDefer(t *testing.T) {
	_defer.DeferInLoop()
}

func TestIoClose(t *testing.T) {
	resp, err := _defer.IoClose("http://49.235.91.10:9200/")
	if err != nil {
		log.Fatalln(err.Error())
	}
	data, _ := ioutil.ReadAll(resp.Body)
	log.Printf("response status: %d, response: %v", resp.StatusCode, data)
}

func TestOsExit(t *testing.T) {
	_defer.OsExit()
}
