package whsvc

import (
	"server/common/levellog"
	"server/dao/whdao"
	"server/models"
)

func Update(warehouse models.Warehouse) error {
	err := whdao.Update(warehouse)
	if err != nil {
		levellog.Service("更新失败")
		return err
	}
	return nil
}
