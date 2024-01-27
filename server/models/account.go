package models

import (
	"database/sql"
	"time"
)

type Account struct {
	Username  string       `json:"username" gorm:"primarykey"`
	Password  string       `json:"password"`
	Salt      string       `json:"salt"`
	Promise   int          `json:"promise"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index"`
}

func (a Account) TableName() string {
	return "account"
}
