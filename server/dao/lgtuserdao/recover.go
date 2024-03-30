package lgtuserdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// RecoverUser 恢复用户信息
func RecoverUser(username string) (user models.User, err error) {
	db := dao.GetMySQL()
	tx := db.Begin()
	err = tx.Model(&models.LogoutUser{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("无法从注销用户表中查找到该用户，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.User{}).Create(user).Error
	if err != nil {
		levellog.Dao("无法将用户信息添加到user表中")
		tx.Rollback()
		return models.User{}, err
	}
	var ac models.Account
	err = tx.Unscoped().Where("username = ?", username).Take(&ac).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("无法从account表中查找到username = %s的记录", username))
		tx.Rollback()
		return models.User{}, err
	}
	ac.DeletedAt = nil
	err = tx.Save(&ac).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("无法在account表中将账号恢复，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.LogoutUser{}).Where("username = ?", username).Delete(&models.LogoutUser{}).Error
	if err != nil {
		levellog.Dao("无法将用户信息从logout_user表中删除")
		tx.Rollback()
		return models.User{}, err
	}
	tx.Commit()
	return user, nil
}
