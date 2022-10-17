package test

import (
	"testing"

	"cn.edu.ntu.awesome/v2/syntax/pointer"
)

func TestAllocate(t *testing.T) {
	pointer.DefineThenAllocate()
}
