package models

type Employee struct {
	ID       uint   `json:"ID" gorm:"primarykey" form:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Position string `json:"position" form:"position"`
	Salary   int    `json:"salary" form:"salary"`
}

func (e Employee) TableName() string {
	return "employee"
}
