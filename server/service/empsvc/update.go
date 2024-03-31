package empsvc

import (
	"fmt"
	"server/common/levellog"
	"server/dao/empdao"
	"server/models"
)

// Update 更新员工信息
func Update(employee models.Employee) error {
	emp, err := empdao.Update(employee)
	if err != nil {
		levellog.Service(fmt.Sprintf("更新失败， emp: %v", emp))
		return err
	}
	return nil
}
