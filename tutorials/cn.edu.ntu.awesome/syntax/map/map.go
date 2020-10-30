package _map

import (
	"fmt"
	"log"
	rand "math/rand"
	"sort"
	"time"
)

var nameAndAgeMap map[string]int

func init() {
	nameAndAgeMap = make(map[string]int, 10)
	nameAndAgeMap["zack"] = 100
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
}

func ApiUsage() {

	// 1. value, ok := map[key]
	v, ok := nameAndAgeMap["zack"]
	if ok {
		log.Println(v)
	} else {
		log.Println("not find")
	}

	// 2. traverse map
	for k, v := range nameAndAgeMap {
		log.Println(k, v)
	}
	// traverse map keys
	for k := range nameAndAgeMap {
		fmt.Println(k)
	}

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
