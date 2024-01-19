package slice

import (
	"log"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSize(t *testing.T) {
	var es []int
	assert.Nil(t, es)
	assert.NotEqualValues(t, unsafe.Sizeof(es), 0)

	var es2 = []int{}
	assert.NotNil(t, es2)
	assert.NotEqualValues(t, unsafe.Sizeof(es2), 0)

	var v struct{}
	assert.NotNil(t, v)
	assert.EqualValues(t, unsafe.Sizeof(v), 0)
}

func TestSliceCreate(t *testing.T) {
	slice1 := CreateSlice()
	log.Print(slice1)

	sliceCopy := slice1
	sliceCopy[0] = 100

	log.Print(slice1)
	log.Print(&sliceCopy[0] == &slice1[0])

}

func TestSliceInits(t *testing.T) {
	InitSlices()
}
