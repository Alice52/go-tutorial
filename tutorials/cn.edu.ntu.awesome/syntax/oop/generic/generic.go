package generic

import (
	"fmt"
)

// type Set[T Animal | Dog] interface {
type Set[T any] interface {
	Put(val T)
	All() []T
}

type HashSet[T any] struct {
	Elements []T
}

func (t *HashSet[T]) Put(val T) {
	fmt.Printf("call method of put(%v)\n", val)
	t.Elements = append(t.Elements, val)
}

func (t *HashSet[T]) All() []T {
	fmt.Printf("call method of all(): %v\n", t.Elements)
	return t.Elements
}
