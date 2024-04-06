package usersvc

import (
	"server/common/levellog"
	"server/dao/pwdmg"
	"server/dao/userdao"
	"server/request"
	"server/utils/pwdtool"
)

// ForgetPassword 用户忘记密码
func ForgetPassword(req request.ForgetPwdReq) (err error) {
	// 校验数据库是否有该用户并且校验绑定的号码是否正确
	_, err = userdao.SelectByUsernameAndPhone(req.Username, req.Phone)
	if err != nil {
		levellog.Service("用户绑定电话错误")
		return err
	}
	// 对新密码加密
	pwd, err := pwdtool.GetPwd(req.Password)
	if err != nil {
		levellog.Service("新密码加密失败")
		return err
	}
	// 更新新密码
	err = pwdmg.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levellog.Service("修改密码错误")
		return err
	}
	return nil
}
