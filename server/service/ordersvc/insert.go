package ordersvc

import (
	"errors"
	"server/common/levellog"
	"server/dao/fruitdao"
	"server/dao/orderdao"
	"server/dao/userdao"
	"server/dao/whdao"
	"server/models"
)

// Insert 用于插入订单
func Insert(order models.Order, username string) error {
	// 检查下单的用户是否存在，如果不存在，则返回错误
	uId, err := userdao.GetUId(username)
	if err != nil {
		levellog.Service("获取用户id失败")
		return errors.New("用户不存在")
	}
	order.BuyerID = int64(uId)

	// 检查下单的用户是否存在，如果不存在，则返回错误
	if has := userdao.ExistsById(int64(uId)); !has {
		return errors.New("用户不存在")
	}

	// 检查仓库id是否存在，并且仓库是否有足够的库存
	if has := whdao.Has(order.WarehouseID); !has {
		return errors.New("仓库不存在")
	}

	// 检查订单中的水果信息是否存在，并且库存是否充足	1吨果蔬大于需要4~5立方米的空间
	// 如果水果不存在或者库存不足，则返回错误
	if has := fruitdao.Has(order.CommodityID); !has {
		return errors.New("水果不存在")
	}

	// 插入订单到数据库
	return orderdao.Insert(order)
}
