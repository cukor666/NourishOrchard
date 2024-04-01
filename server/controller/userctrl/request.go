package userctrl

import (
	"server/models"
	"time"
)

type UpdateUserReq struct {
	Username string    `json:"username" binding:"required,username"`
	Name     string    `json:"name" binding:"omitempty,min=2,max=20"`
	Gender   string    `json:"gender" binding:"omitempty,gender"`
	Phone    string    `json:"phone" binding:"omitempty,phone"`
	Address  string    `json:"address" binding:"omitempty,gte=3,lte=100"`
	Birthday time.Time `json:"birthday" binding:"omitempty,birthday"`
}

func (uur *UpdateUserReq) ToUser() models.User {
	return models.User{
		ID:       0,
		Username: uur.Username,
		Name:     uur.Name,
		Gender:   uur.Gender,
		Phone:    uur.Phone,
		Address:  uur.Address,
		Birthday: uur.Birthday,
	}
}
