package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint         `json:"ID" gorm:"primarykey"`
	Name      string       `json:"name" gorm:"size:20;unique;not null"`
	NickName  string       `json:"nickname" gorm:"size:20;not null"`
	Password  string       `json:"password" gorm:"size:20;default:123456"`
	Promise   int8         `json:"promise" gorm:"not null;default:1"`
	Gender    string       `json:"gender" gorm:"size:4;default:男"`
	Birthday  *time.Time   `json:"birthday" gorm:"null"`
	Phone     string       `json:"phone" gorm:"size:20;not null;"`
	Address   string       `json:"address" gorm:"not null"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index"`
}

// TableName 设置表名
func (u User) TableName() string {
	return "users"
}
