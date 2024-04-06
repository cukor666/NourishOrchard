package request

import "time"

type RegisterRequest struct {
	Name     string    `json:"name" binding:"required,min=3,max=20"`
	Password string    `json:"password" binding:"required,min=3,max=20"`
	Gender   string    `json:"gender" binding:"required,gender"`
	Phone    string    `json:"phone" binding:"required,phone"`
	Address  string    `json:"address" binding:"required,min=2"`
	Birthday time.Time `json:"birthday" binding:"required,birthday"`
}
