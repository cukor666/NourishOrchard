package lgtusersvc

import (
	"fmt"
	"server/common/levellog"
	"server/dao/lgtuserdao"
)

// RemoveUser 彻底删除用户
func RemoveUser(id uint, username string) error {
	user, err := lgtuserdao.RemoveUser(id, username)
	if err != nil {
		levellog.Service(fmt.Sprintf("删除用户失败，user = %v", user))
		return err
	}
	levellog.Service(fmt.Sprintf("删除用户成功，user = %v", user))
	return nil
}
