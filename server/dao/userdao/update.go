package userdao

import (
	"server/dao"
	"server/models"
)

// Update 更新用户信息
func Update(user models.User) (models.User, error) {
	db := dao.GetMySQL()
	err := db.Model(&user).Omit("id", "username").
		Where("username = ?", user.Username).Updates(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
