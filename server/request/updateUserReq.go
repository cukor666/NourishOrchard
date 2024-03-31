package request

import (
	"server/models"
	"time"
)

type UpdateUserReq struct {
	Username string    `json:"username" binding:"required,username"`
	Name     string    `json:"name,omitempty"`
	Gender   string    `json:"gender,omitempty" binding:"gender"`
	Phone    string    `json:"phone,omitempty" binding:"phone"`
	Address  string    `json:"address,omitempty" binding:"gte=3,lte=100"`
	Birthday time.Time `json:"birthday,omitempty" binding:"birthday"`
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
