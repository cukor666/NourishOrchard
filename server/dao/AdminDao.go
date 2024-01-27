package dao

import "server/models"

type AdminDao struct{}

func (a AdminDao) GetUId(username string) (id uint, err error) {
	err = mysqlDB.Table(models.Admin{}.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}
