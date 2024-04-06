package ivtsvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/ivtdao"
	"server/models"
)

func List(p simpletool.Page, inventory models.Inventory) ([]models.Inventory, int64, error) {
	result, total, err := ivtdao.List(p, inventory)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}
