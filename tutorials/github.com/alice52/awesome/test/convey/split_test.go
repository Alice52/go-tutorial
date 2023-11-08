package convey

import (
	"testing"
)

// 测试集的 Setup 与 Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("executing setup")
	return func(t *testing.T) {
		t.Log("executing teardown")
	}
}
