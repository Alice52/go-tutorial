package test

import (
	"cn.edu.ntu.awesome/v0/syntax/slice"
	"log"
	"testing"
)

func TestSliceCreate(t *testing.T) {
	slice := slice.CreateSlice()
	log.Print(slice)

	sliceCopy := slice
	sliceCopy[0] = 100

	log.Print(slice)
	log.Print(&sliceCopy[0] == &slice[0])

}

func TestSliceInits(t *testing.T) {
	slice.InitSlices()
}
