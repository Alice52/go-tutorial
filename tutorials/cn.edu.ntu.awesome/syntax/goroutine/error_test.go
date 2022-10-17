package goroutine

import "testing"

func TestRecoverError(t *testing.T) {
	DeferRecoverError()
}

func TestGroupConError(t *testing.T) {
	GroupConError()
}
