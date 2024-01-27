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
