package regctrl

import (
	"server/models"
	mc "server/models/code"
	"time"
)

type registerRequest struct {
	Name     string    `json:"name" binding:"required,min=2,max=20"`
	Password string    `json:"password" binding:"required,min=3,max=20"`
	Gender   string    `json:"gender" binding:"required,gender"`
	Phone    string    `json:"phone" binding:"required,phone"`
	Address  string    `json:"address" binding:"required,min=2"`
	Birthday time.Time `json:"birthday" binding:"required,birthday"`
}

func (r registerRequest) Deconstruct() (models.Account, models.User) {
	return models.Account{
			Password: r.Password,
			Promise:  mc.USER, // 该接口只提供用户注册，管理员和员工的注册由内部自己完成
		}, models.User{
			Name:     r.Name,
			Gender:   r.Gender,
			Phone:    r.Phone,
			Address:  r.Address,
			Birthday: r.Birthday,
		}
}
