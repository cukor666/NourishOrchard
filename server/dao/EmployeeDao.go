package dao

import (
	"fmt"
	"server/common/simpletool"
	"server/models"
	mc "server/models/code"
)

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
		mysqlDB.Model(&models.EmployeeStatus{}).Select("id").Where("status = ?", mc.Employed)).
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

// SelectByUsernameAndPhone 通过账号和电话查找员工信息
func (e EmployeeDao) SelectByUsernameAndPhone(username, phone string) (employee models.Employee, err error) {
	err = mysqlDB.Model(&models.Employee{}).Where("username = ? AND phone = ?", username, phone).Take(&employee).Error
	if err != nil {
		levelLog(fmt.Sprintf("查找员工失败, employee = %v", employee))
		return models.Employee{}, err
	}
	return employee, nil
}

// Promotion 晋升管理员
func (e EmployeeDao) Promotion(username, name, email, mark string) error {
	tx := mysqlDB.Begin()
	a := models.Admin{
		Username: username,
		Name:     name,
		Email:    email,
	}

	// 更新account表中的对应账号的权限
	err := tx.Model(&models.Account{}).Where("username = ? AND promise = ?", username, mc.EMPLOYEE).
		Update("promise", mc.ADMIN).Error

	if err != nil {
		str := fmt.Sprintf("从%s表中更新username = %s的权限失败", models.Account{}.TableName(), username)
		levelLog(str)
		tx.Rollback()
		return err
	}

	err = tx.Model(&models.Admin{}).Create(&a).Error
	if err != nil {
		levelLog(fmt.Sprintf("将%v插入到%s表中失败", a, a.TableName()))
		tx.Rollback()
		return err
	}

	// 修改employee_status表的状态
	err = tx.Model(&models.EmployeeStatus{}).Where("username = ?", username).Updates(map[string]interface{}{
		"status": mc.TransferFromAPosition,
		"mark":   mark,
	}).Error

	if err != nil {
		levelLog(fmt.Sprintf("在%s表中修改username = %s失败", models.EmployeeStatus{}.TableName(), username))
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
