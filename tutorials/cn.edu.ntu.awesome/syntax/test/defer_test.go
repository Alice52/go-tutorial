package test

import (
	_defer "cn.edu.ntu.awesome/v0/syntax/defer"
	"io/ioutil"
	"log"
	"testing"
)

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
