package mysql

import (
	"database/sql"
)

// 记录用户浏览产品信息:
//
//	操作views和product_viewers两张表
func recordStats(db *sql.DB, userID, productID int64) (err error) {
	// 1. 开启事务
	tx, err := db.Begin()
	if err != nil {
		return
	}

	// 4. commit/rollback
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	// 2. 更新products表
	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}

	// 3. product_viewers表中插入一条数据
	if _, err = tx.Exec(
		"INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)",
		userID, productID); err != nil {
		return
	}
	return
}
