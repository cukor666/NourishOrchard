package empdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// SelectByUsername 根据账户名查询员工信息
func SelectByUsername(username string) (emp models.Employee, err error) {
	db := dao.GetMySQL()
	err = db.Model(&emp).Where("username = ?", username).Take(&emp).Error
	if err != nil {
		return models.Employee{}, err
	}
	return emp, nil
}

// SelectByUsernameAndPhone 通过账号和电话查找员工信息
func SelectByUsernameAndPhone(username, phone string) (employee models.Employee, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.Employee{}).Where("username = ? AND phone = ?", username, phone).Take(&employee).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("查找员工失败, employee = %v", employee))
		return models.Employee{}, err
	}
	return employee, nil
}
