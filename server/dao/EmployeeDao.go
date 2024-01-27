package dao

import "server/models"

type EmployeeDao struct{}

func (e EmployeeDao) GetUId(username string) (id uint, err error) {
	err = mysqlDB.Table(models.Employee{}.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}
