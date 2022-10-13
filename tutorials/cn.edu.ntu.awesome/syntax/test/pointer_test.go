package test

import (
	"testing"

	"cn.edu.ntu.awesome/v0/syntax/pointer"
)

func TestAllocate(t *testing.T) {
	pointer.DefineThenAllocate()
}
