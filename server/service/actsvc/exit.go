package actsvc

import (
	"server/common/levellog"
	"server/dao/actdao"
)

// Exit 账号退出
func Exit(username string) error {
	err := actdao.Exit(username)
	if err != nil {
		levellog.Service("退出失败")
		return err
	}
	return nil
}
