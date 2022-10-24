package json

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedTime time.Time `json:"created_time"`
}

type tempOrder Order // 定义与Order字段一致的新类型
type order_alias struct {
	CreatedTime string `json:"created_time"`
	*tempOrder         // 避免直接嵌套Order进入死循环
}

const layout = "2006-01-02 15:04:05"

// MarshalJSON 为 Order 类型实现自定义的 MarshalJSON 方法
func (o *Order) MarshalJSON() ([]byte, error) {

	oa := order_alias{
		CreatedTime: o.CreatedTime.Format(layout),
		tempOrder:   (*tempOrder)(o),
	}

	return json.Marshal(oa)
}

// UnmarshalJSON 为 Order 类型实现自定义的 UnmarshalJSON 方法
func (o *Order) UnmarshalJSON(data []byte) error {
	oa := order_alias{
		tempOrder: (*tempOrder)(o),
	}

	if err := json.Unmarshal(data, &oa); err != nil {
		return err
	}
	var err error
	o.CreatedTime, err = time.Parse(layout, oa.CreatedTime)
	if err != nil {
		return err
	}
	return nil
}

// 自定义序列化方法
func ParserUsage() {
	o1 := Order{
		ID:          123456,
		Title:       "《七米的Go学习笔记》",
		CreatedTime: time.Now(),
	}
	// 通过自定义的MarshalJSON方法实现struct -> json string
	b, err := json.Marshal(&o1)
	if err != nil {
		fmt.Printf("json.Marshal o1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
	// 通过自定义的UnmarshalJSON方法实现json string -> struct
	jsonStr := `{"created_time":"2020-04-05 10:18:20","id":123456,"title":"《七米的Go学习笔记》"}`
	var o2 Order
	if err := json.Unmarshal([]byte(jsonStr), &o2); err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("o2:%#v\n", o2)
}
