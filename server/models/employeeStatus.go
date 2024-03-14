package models

type EmployeeStatus struct {
	ID       uint   `json:"id" gorm:"primarykey" form:"id"`
	Username string `json:"username" form:"username"`
	Status   int8   `json:"status" form:"status"`
	Mark     string `json:"mark" form:"mark"`
}

func (e EmployeeStatus) TableName() string {
	return "employee_status"
}
