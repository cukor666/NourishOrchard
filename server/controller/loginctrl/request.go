package loginctrl

import (
	"server/models"
	"server/utils/promisetool"
)

type loginRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,min=6,max=100"`
	Promise  string `json:"promise" binding:"required,promise"`
}

func (r loginRequest) toAccount() models.Account {
	return models.Account{
		Username: r.Username,
		Password: r.Password,
		Promise:  promisetool.ToInt(r.Promise),
	}
}
