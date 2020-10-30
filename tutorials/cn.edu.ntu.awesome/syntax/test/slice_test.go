package test

import (
	_slice "cn.edu.ntu.awesome/v0/syntax/slice"
	"log"
	"testing"
)

func TestSliceCreate(t *testing.T) {
	slice1 := _slice.CreateSlice()
	log.Print(slice1)

	sliceCopy := slice1
	sliceCopy[0] = 100

	log.Print(slice1)
	log.Print(&sliceCopy[0] == &slice1[0])

}

func TestSliceInits(t *testing.T) {
	_slice.InitSlices()
}