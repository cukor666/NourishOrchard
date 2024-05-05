package fruitcmdysvc

import (
	"server/common/levellog"
	"server/dao/fruitcmdydao"
	"server/models"
)

func List() (result []models.FruitCommodity, total int64, err error) {
	result, total, err = fruitcmdydao.List()
	if err != nil {
		levellog.Service("查询失败")
		return
	}
	return
}
