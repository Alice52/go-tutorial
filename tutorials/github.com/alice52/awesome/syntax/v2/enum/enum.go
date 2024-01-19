package enum

import "fmt"

type Enum int

const (
	Sunday Enum = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func NewEnum(enum Enum) (Enum, error) {
	if enum < Sunday || enum > Saturday {
		return -1, fmt.Errorf("invalid enum value")
	}

	return enum, nil
}

func Value(enum Enum) {
}
