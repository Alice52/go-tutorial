package json

import (
	"encoding/json"
	"fmt"
	"time"
)

// Deprecated: xx
type Post struct {
	CreateTime time.Time `json:"create_time"`
}

func ParseDate() {
	p1 := Post{CreateTime: time.Now()}
	b, _ := json.Marshal(p1)
	fmt.Printf("str:%s\n", b)

	jsonStr := `{"create_time":"2020-04-05 12:25:42"}`
	var p2 Post
	if err := json.Unmarshal([]byte(jsonStr), &p2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}
