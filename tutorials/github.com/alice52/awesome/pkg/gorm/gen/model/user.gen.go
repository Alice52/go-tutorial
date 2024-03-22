// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID        int64          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true;comment:ID" json:"id"`             // ID
	Name      string         `gorm:"column:name;type:varchar(20);not null;comment:用户名" json:"name"`                              // 用户名
	Age       int64          `gorm:"column:age;type:tinyint unsigned;not null;comment:年龄" json:"age"`                            // 年龄
	Balance   float64        `gorm:"column:balance;type:decimal(11,2) unsigned;not null;default:0.00;comment:余额" json:"balance"` // 余额
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;comment:更新时间" json:"updated_at"`                    // 更新时间
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;comment:创建时间" json:"created_at"`                    // 创建时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                             // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
