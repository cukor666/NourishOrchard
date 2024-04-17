package ordersvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/orderdao"
	"server/models"
)

func List(p simpletool.Page, order models.Order) ([]models.Order, int64, error) {
	result, total, err := orderdao.List(p, order)
	if err != nil {
		levellog.Service("查询订单列表失败")
		return nil, 0, err
	}
	return result, total, nil
}
