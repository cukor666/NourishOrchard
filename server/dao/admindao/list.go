package admindao

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// ListWithPage 查询管理员接口
func ListWithPage(p simpletool.Page) (result []models.Admin, total int64, err error) {
	db := dao.GetMySQL()
	tx := db.Model(&models.Admin{}).Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levellog.Dao("查询管理员列表失败")
		return nil, 0, err
	}
	return result, total, nil
}
