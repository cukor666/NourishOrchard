package adminsvc

import (
	"errors"
	"server/common/levellog"
	"server/dao/admindao"
	"server/dao/pwdmg"
	"server/request"
	"server/utils/pwdtool"
)

func ForgetPassword(req request.ForgetPwdReq) error {
	// 查询管理员是否存在
	_, err := admindao.SelectByUsernameAndEmail(req.Username, req.Email)
	if err != nil {
		levellog.Service("查询管理员失败, 无该管理员")
		return errors.New("查询管理员失败, 无该管理员")
	}
	// 对新密码加密
	pwd, err := pwdtool.GetPwd(req.Password)
	if err != nil {
		levellog.Service("新密码加密失败")
		return err
	}
	// 修改密码
	err = pwdmg.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levellog.Service("修改密码错误")
		return err
	}
	return nil
}
