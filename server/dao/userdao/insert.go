package userdao

import (
	"server/dao"
	"server/models"
)

func Insert(user models.User) (uint, bool) {
	db := dao.GetMySQL()
	tx := db.Create(&user)
	if tx.Error != nil {
		return 0, false
	}
	return user.ID, true
}
