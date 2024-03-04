package dao

import "server/models"

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
		return user, nil
	}
	return models.User{}, err
}
