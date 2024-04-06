package empdao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
	mc "server/models/code"
)

// ListWithPage 查询员工列表，带有分页
func ListWithPage(p simpletool.Page, employee models.Employee) (result []models.Employee, total int64, err error) {
	id, username, name, phone, position, salary := employee.SetZero()
	employee.ID = id
	employee.Salary = salary
	db := dao.GetMySQL()

	tx := db.Model(&employee).Where("id IN (?)",
		db.Model(&models.EmployeeStatus{}).Select("id").Where("status = ?", mc.Employed)).
		Where(&employee).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Where("position LIKE ?", fmt.Sprintf("%%%s%%", position)).
		Count(&total)
	levellog.Dao(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levellog.Dao("查询员工信息失败")
		return nil, 0, err
	}
	return result, total, nil
}
