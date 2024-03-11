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
func (u UserDao) ListWithPage(p simpletool.Page) (result []models.User, total int64, err error) {
	tx := mysqlDB.Model(&models.User{}).Count(&total)
	levelLog(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
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
