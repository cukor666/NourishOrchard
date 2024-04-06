package usersvc

import (
	"server/common/levellog"
	"server/dao/userdao"
	"server/models"
)

// DeleteUser 根据账号删除用户信息
func DeleteUser(username string) (models.User, error) {
	user, err := userdao.DeleteByUsername(username)
	if err != nil {
		levellog.Service("用户删除失败")
		return models.User{}, err
	}
	return user, nil
}
