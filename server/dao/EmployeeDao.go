package dao

import "server/models"

type EmployeeDao struct{}

func (e EmployeeDao) GetUId(username string) (id uint, err error) {
	err = mysqlDB.Table(models.Employee{}.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}

// SelectByUsername 根据账户名查询员工信息
func (e EmployeeDao) SelectByUsername(username string) (emp models.Employee, err error) {
	err = mysqlDB.Model(&emp).Where("username = ?", username).Take(&emp).Error
	if err != nil {
		return models.Employee{}, err
	}
	return emp, nil
}

// Update 更新员工信息
func (e EmployeeDao) Update(employee models.Employee) (models.Employee, error) {
	err := mysqlDB.Model(&employee).Omit("id", "username").
		Where("username = ?", employee.Username).Updates(&employee).Error
	if err != nil {
		return models.Employee{}, err
	}
	return employee, nil
}
