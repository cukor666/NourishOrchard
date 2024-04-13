package empsvc

import (
	"server/common/levellog"
	"server/dao/empdao"
	"server/dao/pwdmg"
	"server/models"
	"server/utils/pwdtool"
)

// ForgetPassword 忘记密码
func ForgetPassword(req models.Employee, password string) error {
	_, err := empdao.SelectByUsernameAndPhone(req.Username, req.Phone)
	if err != nil {
		levellog.Service("查询员工失败")
		return err
	}
	// 对新密码加密
	pwd, err := pwdtool.GetPwd(password)
	if err != nil {
		levellog.Service("新密码加密失败")
		return err
	}
	err = pwdmg.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levellog.Service("修改密码错误")
		return err
	}
	return nil
}
