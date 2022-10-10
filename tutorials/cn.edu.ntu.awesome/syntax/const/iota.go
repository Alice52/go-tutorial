package _const

import (
	"fmt"
	"log"
)

// all followed var will be init as previous
const (
	a1 = 10
	a2
	a3
)

func init() {
	fmt.Printf("a1: %v\n", a1)
	fmt.Printf("a2: %v\n", a2)
	fmt.Printf("a3: %v\n", a3)
}

const (
	a, b = iota + 1, iota + 2 // 1,2
	c, d                      // 2,3
	e, f                      // 3,4
)

const (
	n1 = iota // 0
	n2 = 100  // 100
	_
	n3 // 3
	n4 // 4
)
const n5 = iota // 0

func IotaValue() {
	log.Printf("a: %d, b: %d, c: %d, d: %d, e: %d, f: %d", a, b, c, d, e, f)

	log.Printf("n1: %d, n2: %d, n3: %d, n4: %d", n1, n2, n3, n4)

	log.Printf("n5: %d", n5)
}
