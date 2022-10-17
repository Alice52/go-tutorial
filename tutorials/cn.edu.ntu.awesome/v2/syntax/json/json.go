package json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type Student struct {
	ID     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []*Student
}

func JsonDecoder() {
	const jsonStream = `
		{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
	`

	type Message struct {
		Name, Text string
	}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

// JSON序列化
func Serialize() {

	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化: 结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
}

// JSON反序列化:
func DeSerialize() {
	//JSON反序列化: JSON格式的字符串-->结构体
	str := `
	{
		"Title": "101",
		"Students": [
			{
				"ID": 0,
				"Gender": "男",
				"Name": "stu00"
			},
			{
				"ID": 1,
				"Gender": "男",
				"Name": "stu01"
			}
		]
	}`
	c1 := &Class{}
	err := json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
