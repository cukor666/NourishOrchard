package actsvc

import (
	"errors"
	"fmt"
	"server/common/levellog"
	"server/dao/pwdmg"
	"server/utils/pwdtool"
)

// ChangePassword 修改密码
func ChangePassword(username, oldPassword, newPassword string) error {
	// 从数据库中获取原密码
	password, err := pwdmg.GetPassword(username)
	if err != nil {
		levellog.Service("获取密码失败")
		return err
	}
	// 将前端传递的密码和数据库中的密码进行比较
	if !pwdtool.PwdOK(password, oldPassword) {
		levellog.Service("前端传递密码与数据库中原密码不一致，拒绝修改密码请求")
		return errors.New("前端传递密码与数据库中密码不一致")
	}
	// 对新密码进行加密
	newPwd, err := pwdtool.GetPwd(newPassword)
	if err != nil {
		levellog.Service(fmt.Sprintf("对密码加密失败, newPassword = %s", newPassword))
		return err
	}
	err = pwdmg.ChangePassword(username, string(newPwd))
	if err != nil {
		levellog.Service("修改密码失败")
		return err
	}
	return nil
}
