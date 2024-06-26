package whdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Has 判断id是否存在
func Has(id int64) bool {
	db := dao.GetMySQL()
	cnt := int64(0)
	if err := db.Model(&models.Warehouse{}).Where("id = ?", id).Count(&cnt).Error; err != nil {
		levellog.Dao("系统错误查询失败")
		return false
	}
	if cnt == 0 {
		levellog.Dao(fmt.Sprintf("%s表中不存在id = %d的数据", (&models.Warehouse{}).TableName(), id))
		return false
	}
	return true
}
