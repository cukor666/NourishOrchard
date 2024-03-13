package dao

import (
	"fmt"
	"server/common/simpletool"
	"server/models"
)

func (u UserDao) Insert(user models.User) (uint, bool) {
	tx := mysqlDB.Create(&user)
	if tx.Error != nil {
		return 0, false
	}
	return user.ID, true
}

func (u UserDao) GetUId(username string) (id uint, err error) {
	err = mysqlDB.Table(models.User{}.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}

// SelectByUsername 通过username查询查询用户
func (u UserDao) SelectByUsername(username string) (user models.User, err error) {
	err = mysqlDB.Model(&user).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Update 更新用户信息
func (u UserDao) Update(user models.User) (models.User, error) {
	err := mysqlDB.Model(&user).Omit("id", "username").
		Where("username = ?", user.Username).Updates(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// ListWithPage 查询普通用户列表，分页查询
func (u UserDao) ListWithPage(p simpletool.Page, user models.User) (result []models.User, total int64, err error) {
	id, username, name, gender, phone, address, _ := user.SetZero()
	user.ID = id
	tx := mysqlDB.Model(&user).Where(&user).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("gender LIKE ?", fmt.Sprintf("%%%s%%", gender)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Where("address LIKE ?", fmt.Sprintf("%%%s%%", address)).
		Count(&total)
	levelLog(fmt.Sprintf("total = %d", total))
	err = tx.Model(&user).Where(&user).
		Where("username LIKE ?", fmt.Sprintf("%%%s%%", username)).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("gender LIKE ?", fmt.Sprintf("%%%s%%", gender)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Where("address LIKE ?", fmt.Sprintf("%%%s%%", address)).
		Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询用户列表失败")
		return nil, 0, err
	}
	return
}

// DeleteByUsername 根据账号删除用户信息
func (u UserDao) DeleteByUsername(username string) (user models.User, err error) {
	tx := mysqlDB.Begin()
	err = tx.Model(&models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		levelLog(fmt.Sprintf("无此用户，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.LogoutUser{}).Create(user).Error
	if err != nil {
		levelLog(fmt.Sprintf("将注销用户加入到logout_users表失败，user = %v", user))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.User{}).Where("username = ?", username).Delete(&models.User{}).Error
	if err != nil {
		levelLog(fmt.Sprintf("删除用户失败，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.Account{}).Where("username = ?", username).Delete(&models.Account{}).Error
	if err != nil {
		levelLog(fmt.Sprintf("从account表中删除数据失败，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	tx.Commit()
	return user, err
}

// LogoutListWithPage 查询注销用户列表，带有分页查询
func (u UserDao) LogoutListWithPage(p simpletool.Page) (result []models.LogoutUser, total int64, err error) {
	tx := mysqlDB.Model(&models.LogoutUser{}).Count(&total)
	levelLog(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询注销用户失败")
		return nil, 0, err
	}
	return result, total, nil
}

// RecoverUser 恢复用户信息
func (u UserDao) RecoverUser(username string) (user models.User, err error) {
	tx := mysqlDB.Begin()
	err = tx.Model(&models.LogoutUser{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		levelLog(fmt.Sprintf("无法从注销用户表中查找到该用户，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.User{}).Create(user).Error
	if err != nil {
		levelLog("无法将用户信息添加到user表中")
		tx.Rollback()
		return models.User{}, err
	}
	var ac models.Account
	err = tx.Unscoped().Where("username = ?", username).Take(&ac).Error
	if err != nil {
		levelLog(fmt.Sprintf("无法从account表中查找到username = %s的记录", username))
		tx.Rollback()
		return models.User{}, err
	}
	ac.DeletedAt = nil
	err = tx.Save(&ac).Error
	if err != nil {
		levelLog(fmt.Sprintf("无法在account表中将账号恢复，username = %s", username))
		tx.Rollback()
		return models.User{}, err
	}
	err = tx.Model(&models.LogoutUser{}).Where("username = ?", username).Delete(&models.LogoutUser{}).Error
	if err != nil {
		levelLog("无法将用户信息从logout_user表中删除")
		tx.Rollback()
		return models.User{}, err
	}
	tx.Commit()
	return user, nil
}
