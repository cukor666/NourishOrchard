package models

type Admin struct {
	ID       uint   `json:"ID" gorm:"primarykey"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (a Admin) TableName() string {
	return "admin"
}
