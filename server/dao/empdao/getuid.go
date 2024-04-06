package empdao

import (
	"server/dao"
	"server/models"
)

func GetUId(username string) (id uint, err error) {
	var em models.Employee
	db := dao.GetMySQL()
	err = db.Table(em.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}
