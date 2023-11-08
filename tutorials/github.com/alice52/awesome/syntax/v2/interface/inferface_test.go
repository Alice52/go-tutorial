package _interface

import "testing"

func TestAssertType(t *testing.T) {
	AssertType()
	JustifyType("str")
}

func TestBirdFly(t *testing.T) {

	var bird Flyable = &Bird{
		Name: "bird",
	}

	bird.Fly()
}
