package userdao

import (
	"server/dao"
	"server/models"
)

func GetUId(username string) (id uint, err error) {
	var um models.User
	db := dao.GetMySQL()
	err = db.Table(um.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}
