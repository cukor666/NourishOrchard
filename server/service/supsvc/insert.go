package supsvc

import (
	"server/common/levellog"
	"server/dao/supdao"
	"server/models"
)

// Insert 添加供应商信息
func Insert(supplier models.Supplier) bool {
	ok := supdao.Insert(supplier)
	if !ok {
		levellog.Service("添加供应商信息失败")
		return false
	}
	return true
}
