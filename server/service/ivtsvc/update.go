package ivtsvc

import (
	"server/common/levellog"
	"server/dao/ivtdao"
	"server/models"
)

func Update(inventory models.Inventory) error {
	err := ivtdao.Update(inventory)
	if err != nil {
		levellog.Service("更新失败")
		return err
	}
	return nil
}
