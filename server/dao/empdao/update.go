package empdao

import (
	"server/dao"
	"server/models"
)

// Update 更新员工信息
func Update(employee models.Employee) (models.Employee, error) {
	db := dao.GetMySQL()
	err := db.Model(&employee).Omit("id", "username").
		Where("username = ?", employee.Username).Updates(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}
