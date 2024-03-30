package lgtuserdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// RemoveUser 彻底删除用户
func RemoveUser(id uint, username string) (user models.LogoutUser, err error) {
	db := dao.GetMySQL()
	tx := db.Begin()
	err = tx.Model(&models.LogoutUser{}).
		Where("id = ?", id).Where("username = ?", username).Delete(&user).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("从%s表中删除用户失败, user = %v", user.TableName(), user))
		tx.Rollback()
		return models.LogoutUser{}, err
	}
	err = tx.Model(&models.Account{}).Unscoped().Where("username = ?", username).Delete(&models.Account{}).Error
	if err != nil {
		var am models.Account
		levellog.Dao(fmt.Sprintf("从%s表中删除%s用户失败", am.TableName(), username))
		tx.Rollback()
		return models.LogoutUser{}, err
	}
	tx.Commit()
	return user, nil
}
