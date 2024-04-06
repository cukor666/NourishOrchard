package userdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// DeleteByUsername 根据账号删除用户信息
func DeleteByUsername(username string) (user models.User, err error) {
	db := dao.GetMySQL()
	tx := db.Begin()
	err = tx.Model(&models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("无此用户，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.LogoutUser{}).Create(user).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("将注销用户加入到logout_users表失败，user = %v", user))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.User{}).Where("username = ?", username).Delete(&models.User{}).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("删除用户失败，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.Account{}).Where("username = ?", username).Delete(&models.Account{}).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("从account表中删除数据失败，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	tx.Commit()
	return user, err
}
