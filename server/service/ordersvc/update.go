package ordersvc

import (
	"errors"
	"server/dao/admindao"
	"server/dao/fruitdao"
	"server/dao/orderdao"
	"server/dao/whdao"
	"server/models"
)

// Update 更新订单信息
func Update(order models.Order, username string) error {
	// 进一步校验订单信息，主要是检查数据库中是否存在相关数据
	// 校验操作的管理员id是否存在
	uId, err := admindao.GetUId(username)
	if err != nil {
		return errors.New("管理员不存在")
	}
	order.AdminId = int64(uId)
	if exist, err := admindao.ExistById(order.AdminId); err != nil || !exist {
		return errors.New("管理员不存在")
	}
	// 校验仓库id是否存在
	if has := whdao.Has(order.WarehouseID); !has {
		return errors.New("仓库不存在")
	}
	// 校验商品id是否存在
	if has := fruitdao.Has(order.CommodityID); !has {
		return errors.New("商品不存在")
	}
	// 校验通过完后，更新订单信息
	// 注意：更新订单信息时，需要更新订单状态，如果订单状态发生变化，需要通知相关人员
	return orderdao.Update(order)
}
