package models

import "time"

type LogoutUser struct {
	ID       uint      `json:"id" gorm:"primarykey" form:"id"`
	Username string    `json:"username" form:"username"`
	Name     string    `json:"name" form:"name"`
	Gender   string    `json:"gender" form:"gender"`
	Phone    string    `json:"phone" form:"phone"`
	Address  string    `json:"address" form:"address"`
	Birthday time.Time `json:"birthday" form:"birthday"`
}

func (l LogoutUser) TableName() string {
	return "logout_users"
}
