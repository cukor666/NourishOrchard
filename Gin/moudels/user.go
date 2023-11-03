package moudels

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string    `json:"name" gorm:"size:20;unique;not null"`
	Password string    `json:"password" gorm:"size:20;default:123456"`
	Gender   string    `json:"gender" gorm:"size:4;default:男"`
	Birthday time.Time `json:"birthday" gorm:"null"`
	Phone    string    `json:"phone" gorm:"size:20;not null;"`
	Address  string    `json:"address" gorm:"not null"`
}

func NewUser(name, password, gender, phone, address string) *User {
	user := new(User)
	user.Name = name
	user.Password = password
	user.Gender = gender
	user.Birthday = time.Date(2000, 8, 8, 6, 6, 6, 6, time.UTC)
	user.Phone = phone
	user.Address = address
	return user
}

func (u *User) Show() {
	log.Println("name: ", u.Name)
	log.Println("password: ", u.Password)
	log.Println("gender: ", u.Gender)
	log.Println("birthday: ", u.Birthday)
	log.Println("phone: ", u.Phone)
	log.Println("address: ", u.Address)
}
