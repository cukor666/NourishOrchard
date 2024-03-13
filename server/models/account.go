package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	Username  string          `json:"username" gorm:"primarykey" form:"username"`
	Password  string          `json:"password" form:"password"`
	Promise   int             `json:"promise" form:"promise"`
	CreatedAt time.Time       `json:"createdAt" form:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt" form:"updatedAt"`
	DeletedAt *gorm.DeletedAt `json:"deletedAt" gorm:"index" form:"deletedAt"`
}

func (a Account) TableName() string {
	return "account"
}
