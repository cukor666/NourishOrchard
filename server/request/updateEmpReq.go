package request

import "server/models"

type UpdateEmpReq struct {
	Username string `json:"username" binding:"required,username"`
	Name     string `json:"name,omitempty" binding:"min=2,max=20"`
	Phone    string `json:"phone,omitempty" binding:"phone"`
	Position string `json:"position,omitempty" binding:"empPosition"`
	Salary   int    `json:"salary,omitempty" binding:"gte=3000,lte=30000"`
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
