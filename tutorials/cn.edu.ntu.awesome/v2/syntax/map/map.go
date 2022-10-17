package _map

import (
	"fmt"
	"log"
	rand "math/rand"
	"sort"
	"time"
)

var (
	nameAndAgeMap map[string]int // init as nil
	age           int            // init as 0
)

func init() {
	age = 18
	nameAndAgeMap = make(map[string]int, 10)
	nameAndAgeMap["zack"] = 100
}

func AnyType() {
	any := map[string]interface{}{
		"name": "zack",
		"age":  18,
	}
	fmt.Println(any) // map[age:18 name:zack]
	v, _ := any["name"]
	log.Println(v) // zack

	switch value := v.(type) {
	case string:
		fmt.Printf("string type, and value is %s", value)
	default:
		fmt.Println(value)
	}

	any2 := &map[string]interface{}{
		"name": "zack",
		"age":  18,
	}
	fmt.Println(any2)  // &map[age:18 name:zack]
	fmt.Println(*any2) // map[age:18 name:zack]
	v2, _ := (*any2)["name"]
	log.Println(v2) // zack
}

func DeclareAndInitial() {
	// 1. declare and allocate map
	nameAndAgeMap = make(map[string]int, 10)
	nameAndAgeMap["zack"] = 100
	log.Printf("type of nameAndAgeMap: %T", nameAndAgeMap)

	// 2. create and init
	userInfo := map[string]string{
		"username": "prof.cn",
		"password": "123456",
	}
	fmt.Println(userInfo)

	// 3. get ptr
	ptr := &userInfo
	fmt.Printf("ptr: %v\n", ptr)
}

func Exist() {
	// 1. value, ok := map[key]
	v, ok := nameAndAgeMap["zack"]
	if ok {
		log.Println(v)
	} else {
		log.Println("not find")
	}
}

func Traverse() {
	// 2. traverse map
	for k, v := range nameAndAgeMap {
		log.Println(k, v)
	}
	// traverse map keys
	for k := range nameAndAgeMap {
		fmt.Println(k)
	}
}

func ApiUsage() {

	// 3. delete(map, key)
	delete(nameAndAgeMap, "zack")

	// 4. traverse by specific sequence
	traverseBySequence()

	// 5. use as with slice element
	useAsSlice()

	// 6. value is slice
	valueIsSlice()
}

func valueIsSlice() {
	sliceMap := make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "China"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "BeiJing", "Shanghai")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func useAsSlice() {
	mapSlice := make([]map[string]string, 5)
	log.Printf("mapSlice type: %T", mapSlice)

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v type: %T\n", index, value, value)
	}

	mapSlice[0] = make(map[string]string, 100)
	mapSlice[0]["name"] = "zack"
	mapSlice[0]["age"] = "18"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v type: %T\n", index, value, value)
	}
}

func traverseBySequence() {
	rand.Seed(time.Now().UnixNano())
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		scoreMap[key] = i
	}

	var keys = make([]string, 0, 200)
	for k := range scoreMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		log.Println(v)
	}
}
