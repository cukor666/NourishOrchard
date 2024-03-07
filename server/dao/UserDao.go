package dao

import (
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
func (u UserDao) ListWithPage(p simpletool.Page) (result []models.User, err error) {
	err = mysqlDB.Model(&models.User{}).Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询用户列表失败")
		return nil, err
	}
	return
}

// DeleteByUsername 根据账号删除用户信息
func (u UserDao) DeleteByUsername(username string) error {
	err := mysqlDB.Model(&models.User{}).Delete("username = ?", username).Error
	return err
}

// SelectDeletedUsers 查询被删除的用户，（用于后续恢复用户信息，或进一步清理用户）
func (u UserDao) SelectDeletedUsers() (users []models.User, err error) {
	mysqlDB.Model(&models.User{}).Where("deleted")
	return nil, err
}
