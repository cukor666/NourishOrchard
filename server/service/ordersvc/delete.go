package ordersvc

import (
	"errors"
	"server/common/levellog"
	"server/dao/admindao"
	"server/dao/orderdao"
)

// Delete 删除订单业务
func Delete(id int64, username string) error {
	// 1. 校验权限
	uId, err := admindao.GetUId(username)
	if err != nil {
		levellog.Service("获取管理员ID失败")
		return errors.New("获取管理员ID失败")
	}
	exist, err := admindao.ExistById(int64(uId))
	if err != nil || !exist {
		return errors.New("权限校验失败，管理员不存在")
	}

	// 2. 删除订单
	return orderdao.Delete(id)
}
