package test

import (
	"encoding/json"
	"fmt"
	"github.com/alice52/proxy/common/http"
	"testing"
)

type respResult struct {
	Errno   int    `json:"errno"`
	Errmsg  string `json:"errmsg"`
	IsPrint bool   `json:"isPrint"`
	Data    struct {
		LogId      string   `json:"log_id"`
		ActionRule struct{} `json:"action_rule"`
	} `json:"data"`
}

func TestHttpReq(t *testing.T) {

	body, _ := json.Marshal(map[string]string{
		"key1": "value1",
		"key2": "value2",
	})

	resp, _ := http.DoReq(nil, "POST", "https://ug.baidu.com/mcp/pc/pcsearch", body)
	var rr respResult
	err := json.Unmarshal(resp, &rr)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Printf("%+v\n", rr)

}
