package actctrl

import "server/models"

type forgetPwdReq struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Promise  string `json:"promise" form:"promise" binding:"required,promise"`
	Phone    string `json:"phone" form:"phone" binding:"omitempty,phone"`
	Email    string `json:"email" form:"email" binding:"omitempty,email"`
	Code     string `json:"code" form:"code" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=20"`
}

func (r forgetPwdReq) toUser() models.User {
	return models.User{
		Username: r.Username,
		Phone:    r.Phone,
	}
}

func (r forgetPwdReq) toEmp() models.Employee {
	return models.Employee{
		Username: r.Username,
		Phone:    r.Phone,
	}
}

func (r forgetPwdReq) toAdmin() models.Admin {
	return models.Admin{
		Username: r.Username,
		Email:    r.Email,
	}
}
