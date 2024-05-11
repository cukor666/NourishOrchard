package ordersvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/orderdao"
	"server/dao/userdao"
	"server/models"
)

func ListForUser(p simpletool.Page, username string) ([]models.Order, int64, error) {
	userId, err := userdao.GetUId(username)
	if err != nil {
		levellog.Service("根据用户名获取用户信息失败")
		return nil, 0, err
	}
	orders, total, err := orderdao.List(p, models.Order{
		BuyerID: int64(userId),
	})
	if err != nil {
		levellog.Service("查询订单列表失败")
		return nil, 0, nil
	}
	return orders, total, nil
}
