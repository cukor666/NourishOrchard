package empdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

func Has(id int64) bool {
	db := dao.GetMySQL()
	cnt := int64(0)
	err := db.Model(&models.Employee{}).Where("id = ?", id).Count(&cnt).Error
	if err != nil {
		levellog.Dao("系统错误查询失败")
		return false
	}
	if cnt == 0 {
		levellog.Dao(fmt.Sprintf("%s表中不存在id = %d的数据", (&models.Employee{}).TableName(), id))
		return false
	}
	return true
}
