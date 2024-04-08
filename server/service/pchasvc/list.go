package pchasvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/pchadao"
	"server/models"
)

func List(p simpletool.Page, purchase models.Purchase) ([]models.Purchase, int64, error) {
	purchases, total, err := pchadao.List(p, purchase)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return purchases, total, nil
}
