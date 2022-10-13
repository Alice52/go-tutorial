package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectType(t *testing.T) {
	var a float32 = 3.14
	ReflectType(a) // type:float32
	var b int64 = 100
	ReflectType(b) // type:int64
}

func TestTypeOrKind(t *testing.T) {
	TypeOrKind()
}

func TestValueOf(t *testing.T) {
	var a float32 = 3.14
	var b int64 = 100
	ReflectValue(a) // type is float32, value is 3.140000
	ReflectValue(b) // type is int64, value is 100

	// 将int类型的原始值转换为reflect.Value类型
	c := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", c) // type c :reflect.Value
}

func TestReflectSetValue(t *testing.T) {

	var a int64 = 100
	v := ReflectSetValue(&a)
	fmt.Printf("v: %v\n", v)
}

func TestReflectStruct(t *testing.T) {

	ReflectStruct()
}
