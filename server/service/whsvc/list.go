package whsvc

import (
	"server/common/simpletool"
	"server/dao/whdao"
	"server/models"
)

func List(p simpletool.Page, w models.Warehouse) ([]models.Warehouse, int64, error) {
	return whdao.List(p, w)
}
