package supdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Insert 添加供应商
func Insert(supplier models.Supplier) bool {
	db := dao.GetMySQL()
	err := db.Model(&supplier).Create(&supplier).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("添加供应商失败，sup: %v", supplier))
		return false
	}
	return true
}
