package lib

// 第三方模块: 未实现
type UserInfo struct {
	Name string
}

func GetInfoByUID(int64) (*UserInfo, error) {
	return &UserInfo{}, nil
}
