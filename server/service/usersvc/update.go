package usersvc

import (
	"fmt"
	"server/common/levellog"
	"server/dao/userdao"
	"server/models"
)

// Update 更新用户信息
func Update(user models.User) error {
	u, err := userdao.Update(user)
	if err != nil {
		levellog.Service(fmt.Sprintf("更新成功, user: %v", u))
	}
	return err
}
