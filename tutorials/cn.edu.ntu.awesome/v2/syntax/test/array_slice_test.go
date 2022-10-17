package test

import "testing"

/**
go test array_slice_test.go -bench . -benchmem -gcflags "-N -l"
BenchmarkArray-4          474877              2392 ns/op               0 B/op          0 allocs/op
BenchmarkSlice-4          365318              3265 ns/op            8192 B/op          1 allocs/op

- conclusion
	+ when array is little, array will have higher performance than slice.
*/
func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}

	return x
}



func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}
