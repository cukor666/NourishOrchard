package models

type Employee struct {
	ID       uint   `json:"id" gorm:"primarykey" form:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Phone    string `json:"phone" form:"phone"`
	Position string `json:"position" form:"position"`
	Salary   int    `json:"salary" form:"salary"`
}

func (e *Employee) TableName() string {
	return "employee"
}

// SetZero 设置字段为零值
func (e *Employee) SetZero() (id uint, username, name, phone, position string, salary int) {
	id, username, name, phone, position, salary = e.ID, e.Username, e.Name, e.Phone, e.Position, e.Salary
	e.ID, e.Username, e.Name, e.Phone, e.Position, e.Salary = 0, "", "", "", "", 0
	return
}
