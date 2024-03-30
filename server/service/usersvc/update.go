package usersvc

import (
	"server/dao/userdao"
	"server/models"
)

// Update 更新用户信息
func Update(user models.User) error {
	_, err := userdao.Update(user)
	return err
}
