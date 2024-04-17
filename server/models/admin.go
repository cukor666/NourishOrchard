package models

type Admin struct {
	ID       uint   `json:"id" gorm:"primarykey" form:"id"`
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
}

func (a *Admin) TableName() string {
	return "admin"
}

// SetZero 设置为初始值，此处可以优化，但是懒得做了
func (a *Admin) SetZero() (id uint, username, name, email string) {
	id, username, name, email = a.ID, a.Username, a.Name, a.Email
	a.ID, a.Username, a.Name, a.Email = 0, "", "", ""
	return
}
