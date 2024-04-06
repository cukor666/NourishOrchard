package lgtusersvc

import (
	"server/common/levellog"
	"server/dao/lgtuserdao"
	"server/models"
)

// RecoverUser 恢复用户信息
func RecoverUser(username string) (models.User, error) {
	user, err := lgtuserdao.RecoverUser(username)
	if err != nil {
		levellog.Service("无法恢复用户信息")
		return models.User{}, err
	}
	return user, nil
}
