package model

type R struct {
	ErrCode int    `json:"code"`
	ErrMsg  string `json:"msg"`
	Url     string `json:"url"`
}
