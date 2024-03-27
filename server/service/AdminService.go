package service

import (
	"errors"
	"regexp"
	"server/common/levellog"
	"server/common/simpletool"
	"server/models"
	"server/request"
	"server/utils"
)

// Info 根据账号获取管理员信息
func (a AdminService) Info(username string) (admin models.Admin, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Admin{}, errors.New("账号不符合系统规范")
	}
	return adminDao.SelectByUsername(username)
}

// Update 更新管理员信息
func (a AdminService) Update(admin models.Admin) error {
	_, err := adminDao.Update(admin)
	return err
}

// ListWithPage 查询管理员列表，分页查询
func (a AdminService) ListWithPage(p simpletool.Page) ([]models.Admin, int64, error) {
	result, total, err := adminDao.ListWithPage(p)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

func (a AdminService) ForgetPassword(req request.ForgetPwdReq) error {
	_, err := adminDao.SelectByUsernameAndEmail(req.Username, req.Email)
	if err != nil {
		levellog.Service("查询管理员失败")
		return err
	}
	// 对新密码加密
	pwd, err := utils.GetPwd(req.Password)
	if err != nil {
		levellog.Service("新密码加密失败")
		return err
	}
	err = accountDao.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levellog.Service("修改密码错误")
		return err
	}
	return nil
}
