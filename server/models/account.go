package models

import (
	"gorm.io/gorm"
	"time"
)

// Account 账户模型
type Account struct {
	Username  string          `json:"username" gorm:"primarykey" form:"username" binding:"required"`
	Password  string          `json:"password" form:"password"`
	Promise   int             `json:"promise" form:"promise"`
	CreatedAt time.Time       `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt" form:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index" form:"deletedAt"`
}

// TableName 数据库表名
func (a *Account) TableName() string {
	return "account"
}
