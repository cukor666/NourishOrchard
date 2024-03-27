package models

import "time"

type User struct {
	ID       uint      `json:"id" gorm:"primarykey" form:"id"`
	Username string    `json:"username" form:"username"`
	Name     string    `json:"name" form:"name"`
	Gender   string    `json:"gender" form:"gender"`
	Phone    string    `json:"phone" form:"phone"`
	Address  string    `json:"address" form:"address"`
	Birthday time.Time `json:"birthday" form:"birthday"`
}

func (u *User) TableName() string {
	return "users"
}

// SetZero 返回对应字段原有的值，并设置字段为零值
func (u *User) SetZero() (id uint, username, name, gender, phone, address string, birthday time.Time) {
	id, username, name, gender, phone, address, birthday = u.ID, u.Username, u.Name, u.Gender, u.Phone, u.Address, u.Birthday
	u.ID, u.Username, u.Name, u.Gender, u.Phone, u.Address, u.Birthday = 0, "", "", "", "", "", time.Time{}
	return
}
