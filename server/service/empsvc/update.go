package empsvc

import (
	"server/dao/empdao"
	"server/models"
)

// Update 更新员工信息
func Update(employee models.Employee) error {
	_, err := empdao.Update(employee)
	return err
}
