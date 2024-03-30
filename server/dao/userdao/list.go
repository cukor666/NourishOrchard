package userdao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// ListWithPage 查询普通用户列表，分页查询
func ListWithPage(p simpletool.Page, user models.User) (result []models.User, total int64, err error) {
	id, username, name, gender, phone, address, _ := user.SetZero()
	user.ID = id
	db := dao.GetMySQL()
	tx := db.Model(&user).Where(&user).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("gender LIKE ?", fmt.Sprintf("%%%s%%", gender)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Where("address LIKE ?", fmt.Sprintf("%%%s%%", address)).
		Count(&total)
	levellog.Dao(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levellog.Dao("查询用户列表失败")
		return nil, 0, err
	}
	return
}
