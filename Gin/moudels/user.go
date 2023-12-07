package moudels

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"name" gorm:"size:20;unique;not null"`
	NickName string     `json:"nickname" gorm:"size:20;not null"`
	Password string     `json:"password" gorm:"size:20;default:123456"`
	Gender   string     `json:"gender" gorm:"size:4;default:男"`
	Birthday *time.Time `json:"birthday" gorm:"null"`
	Phone    string     `json:"phone" gorm:"size:20;not null;"`
	Address  string     `json:"address" gorm:"not null"`
	Promise  int8       `json:"promise" gorm:"not null;default:1"`
}

// TableName 设置表名
func (u User) TableName() string {
	return "users"
}

//func NewUser(name, nickname, password, gender, phone, address string) *User {
//	user := new(User)
//	user.Name = name
//	user.NickName = nickname
//	user.Password = password
//	user.Gender = gender
//	user.Birthday = nil
//	user.Phone = phone
//	user.Address = address
//	user.Promise = 1
//	return user
//}

// Show  展示用户信息
func (u User) Show() {
	log.Println("name: ", u.Name)
	log.Println("nickname: ", u.NickName)
	log.Println("password: ", u.Password)
	log.Println("gender: ", u.Gender)
	log.Println("birthday: ", u.Birthday)
	log.Println("phone: ", u.Phone)
	log.Println("address: ", u.Address)
}
