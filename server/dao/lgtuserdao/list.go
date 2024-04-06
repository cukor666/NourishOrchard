package lgtuserdao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// LogoutListWithPage 查询注销用户列表，带有分页查询
func LogoutListWithPage(p simpletool.Page, logoutUser models.LogoutUser) (result []models.LogoutUser, total int64, err error) {
	id, username, name, gender, phone, address, _ := logoutUser.SetZero()
	logoutUser.ID = id
	db := dao.GetMySQL()
	tx := db.Model(&logoutUser).Where(&logoutUser).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("gender LIKE ?", fmt.Sprintf("%%%s%%", gender)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Where("address LIKE ?", fmt.Sprintf("%%%s%%", address)).
		Count(&total)
	levellog.Dao(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levellog.Dao("查询注销用户失败")
		return nil, 0, err
	}
	return result, total, nil
}
