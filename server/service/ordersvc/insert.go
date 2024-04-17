package ordersvc

import (
	"errors"
	"server/dao/admindao"
	"server/dao/fruitdao"
	"server/dao/whdao"
	"server/models"
)

// Insert 用于插入订单 todo : 这里应该检查库存是否充足，但是暂时先不做
func Insert(order models.Order) error {
	// 先检查是谁操作的订单，然后插入订单到数据库
	// 如果操作订单的管理员不存在，则返回错误
	if exist, err := admindao.ExistById(order.AdminId); err != nil || !exist {
		return errors.New("管理员不存在")
	}

	// 检查仓库id是否存在，并且仓库是否有足够的库存
	if has := whdao.Has(order.WarehouseID); !has {
		return errors.New("仓库不存在")
	}

	// 检查订单中的水果信息是否存在，并且库存是否充足
	// 如果水果不存在或者库存不足，则返回错误
	if has := fruitdao.Has(order.CommodityID); !has {
		return errors.New("水果不存在")
	}
	// todo: 这里应该检查库存是否充足，但是暂时先不做

	// 插入订单到数据库

	return nil
}
