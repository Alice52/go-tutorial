package reflect

import (
	"fmt"
	"reflect"
)

func ReflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type:%v, kind:%v\n", t.Name(), t.Kind())
}

type myInt int64

func TypeOrKind() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	ReflectType(a) // type: kind:ptr
	ReflectType(b) // type:myInt kind:int64
	ReflectType(c) // type:int32 kind:int32

	type person struct {
		name string
		age  int
	}
	var d = person{
		name: "zack",
		age:  18,
	}
	ReflectType(d) // type:person kind:struct

	type book struct{ title string }
	var e = &book{title: "zack"}
	ReflectType(e) // type: kind:ptr
}

// switch x.(type)
func ReflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 反射修改值只能是地址变量
func ReflectSetValue(x interface{}) any {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem() 方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}

	return v
}

// 结构体反射
type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func ReflectStruct() {

	stu1 := student{
		Name:  "小王子",
		Score: 90,
	}

	// 1. 获取类型
	t := reflect.TypeOf(stu1)
	fmt.Printf("type: %v, kind: %v\n", t.Name(), t.Kind()) // student struct

	// 2. 获取值
	v := reflect.ValueOf(stu1)
	fmt.Printf("value: %v\n", v)

	// 3. 获取所有字段: 顺序
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
			field.Name,
			field.Index,
			field.Type,
			field.Tag.Get("json"))
	}

	// 4. 获取指定名称的字段
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n",
			scoreField.Name,
			scoreField.Index,
			scoreField.Type,
			scoreField.Tag.Get("json"))
	}

	// 5. 获取所有的方法
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 6. 执行方法:
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
