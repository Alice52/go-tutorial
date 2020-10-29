package test

import (
	"log"
	"testing"
)

func TestValue(t *testing.T) {
	var age int32
	age++
	log.Println(age)
}
