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

func (l *LogoutUser) TableName() string {
	return "logout_users"
}

// SetZero 返回对应字段原有的值，并设置字段为零值
func (l *LogoutUser) SetZero() (id uint, username, name, gender, phone, address string, birthday time.Time) {
	id, username, name, gender, phone, address, birthday = l.ID, l.Username, l.Name, l.Gender, l.Phone, l.Address, l.Birthday
	l.ID, l.Username, l.Name, l.Gender, l.Phone, l.Address, l.Birthday = 0, "", "", "", "", "", time.Time{}
	return
}
