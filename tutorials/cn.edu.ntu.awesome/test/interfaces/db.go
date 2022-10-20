package interfaces

import "database/sql"

// DB 数据接口
type DB interface {
	Get(key string) (int, error)
}

type MySQL struct {
	sql.DB
}

func (m *MySQL) Get(key string) (int, error) {
	// ...
	return 0, nil
}

// GetFromDB 根据 key 从 DB 查询数据的函数
func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key); err == nil {
		return v
	}
	return -1
}
