package models

import "time"

type LogoutUser struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday"`
}

func (l LogoutUser) TableName() string {
	return "logout_users"
}
