package models

type Employee struct {
	ID       uint   `json:"ID" gorm:"primarykey"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}

func (e Employee) TableName() string {
	return "employee"
}
