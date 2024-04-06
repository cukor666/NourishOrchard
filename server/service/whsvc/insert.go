package whsvc

import (
	"server/common/levellog"
	"server/dao/whdao"
	"server/models"
)

func Insert(warehouse models.Warehouse) error {
	err := whdao.Insert(warehouse)
	if err != nil {
		levellog.Service("添加新仓库失败")
		return err
	}
	return nil
}
