package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	Username  string         `json:"username" gorm:"primarykey"`
	Password  string         `json:"password"`
	Promise   int            `json:"promise"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

func (a Account) TableName() string {
	return "account"
}
