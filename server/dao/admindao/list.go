package admindao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// ListWithPage 查询管理员接口
func ListWithPage(p simpletool.Page, admin models.Admin) (result []models.Admin, total int64, err error) {
	id, username, name, email := admin.SetZero()
	admin.ID = id
	db := dao.GetMySQL()
	tx := db.Model(&models.Admin{}).Where(&admin).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("email LIKE ?", fmt.Sprintf("%%%s%%", email)).
		Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levellog.Dao("查询管理员列表失败")
		return nil, 0, err
	}
	return result, total, nil
}
