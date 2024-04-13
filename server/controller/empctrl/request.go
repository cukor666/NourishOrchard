package empctrl

import "server/models"

type UpdateEmpReq struct {
	Username string `json:"username" binding:"required,username"`
	Name     string `json:"name" binding:"omitempty,min=2,max=20"`
	Phone    string `json:"phone" binding:"omitempty,phone"`
	Position string `json:"position" binding:"omitempty,empPosition"`
	Salary   int    `json:"salary" binding:"omitempty,gte=3000,lte=30000"`
}

func (ueq *UpdateEmpReq) ToEmp() models.Employee {
	return models.Employee{
		ID:       0,
		Username: ueq.Username,
		Name:     ueq.Name,
		Phone:    ueq.Phone,
		Position: ueq.Position,
		Salary:   ueq.Salary,
	}
}

type promoteReq struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=20"`
	Name     string `json:"name" form:"name" binding:"required,min=2,max=20"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Mark     string `json:"mark" form:"mark" binding:"max=100"`
}
