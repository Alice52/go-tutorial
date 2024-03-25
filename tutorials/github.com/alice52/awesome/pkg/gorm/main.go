package main

import (
	"fmt"
	"github.com/alice52/awesome/pkg/gorm/gen/dal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDSN = "root:xxx@@(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

// https://blog.csdn.net/qq_51438138/article/details/129220917
// https://mp.weixin.qq.com/s?__biz=MzU5MjAxMDc1Ng==&mid=2247485002&idx=1&sn=c9c3c6c345bfdccf5f50f954929a43aa
func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	dal.SetDefault(db)
	//查询User年龄为18的
	userDo := dal.User.Where(dal.User.Age.Eq(18))
	fmt.Print(userDo)

	user, err := dal.User.SimpleFindByNameAndAge("zhangqiang", 18)
	fmt.Print(user)
}
