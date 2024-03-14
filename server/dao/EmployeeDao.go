package dao

import (
	"fmt"
	"server/common/simpletool"
	"server/models"
)

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

// ListWithPage 查询员工列表，带有分页
func (e EmployeeDao) ListWithPage(p simpletool.Page, employee models.Employee) (result []models.Employee, total int64, err error) {
	id, username, name, phone, position, salary := employee.SetZero()
	employee.ID = id
	employee.Salary = salary
	tx := mysqlDB.Model(&employee).Where("id IN (?)",
		mysqlDB.Model(&models.EmployeeStatus{}).Select("id").Where("status = ?", 1)).
		Where(&employee).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Where("position LIKE ?", fmt.Sprintf("%%%s%%", position)).
		Count(&total)
	levelLog(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询员工信息失败")
		return nil, 0, err
	}
	return result, total, nil
}
