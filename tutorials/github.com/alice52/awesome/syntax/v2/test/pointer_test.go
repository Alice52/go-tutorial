package test

import (
	"testing"

	"github.com/alice52/awesome/syntax/v2/pointer"
)

func TestAllocate(t *testing.T) {
	pointer.DefineThenAllocate()
}
