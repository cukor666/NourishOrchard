package models

type Admin struct {
	ID       uint   `json:"ID" gorm:"primarykey" form:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
}

func (a Admin) TableName() string {
	return "admin"
}
