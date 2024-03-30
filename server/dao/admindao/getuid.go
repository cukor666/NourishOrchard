package admindao

import (
	"server/dao"
	"server/models"
)

func GetUId(username string) (id uint, err error) {
	var am models.Admin
	db := dao.GetMySQL()
	err = db.Table(am.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}
