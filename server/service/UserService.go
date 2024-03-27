package service

import (
	"errors"
	"fmt"
	"regexp"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/pwdmg"
	"server/models"
	"server/request"
	"server/utils/pwdtool"
)

// Info 根据用户名获取用户信息
func (u UserService) Info(username string) (user models.User, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.User{}, errors.New("用户名不符合系统规范")
	}
	return userDao.SelectByUsername(username)
}

// Update 更新用户信息
func (u UserService) Update(user models.User) error {
	_, err := userDao.Update(user)
	return err
}

// List 访问用户列表，需要管理员或员工权限
func (u UserService) List(p simpletool.Page, user models.User) ([]models.User, int64, error) {
	result, total, err := userDao.ListWithPage(p, user)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// DeleteUser 根据账号删除用户信息
func (u UserService) DeleteUser(username string) (models.User, error) {
	user, err := userDao.DeleteByUsername(username)
	if err != nil {
		levellog.Service("用户删除失败")
		return models.User{}, err
	}
	return user, nil
}

// LogoutList 查询注销用户列表，需要管理员或员工权限
func (u UserService) LogoutList(p simpletool.Page, logoutUser models.LogoutUser) ([]models.LogoutUser, int64, error) {
	result, total, err := userDao.LogoutListWithPage(p, logoutUser)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// RecoverUser 恢复用户信息
func (u UserService) RecoverUser(username string) (models.User, error) {
	user, err := userDao.RecoverUser(username)
	if err != nil {
		levellog.Service("无法恢复用户信息")
		return models.User{}, err
	}
	return user, nil
}

// RemoveUser 彻底删除用户
func (u UserService) RemoveUser(id uint, username string) error {
	user, err := userDao.RemoveUser(id, username)
	if err != nil {
		levellog.Service(fmt.Sprintf("删除用户失败，user = %v", user))
		return err
	}
	levellog.Service(fmt.Sprintf("删除用户成功，user = %v", user))
	return nil
}

// ForgetPassword 用户忘记密码
func (u UserService) ForgetPassword(req request.ForgetPwdReq) (err error) {
	// 校验数据库是否有该用户并且校验绑定的号码是否正确
	_, err = userDao.SelectByUsernameAndPhone(req.Username, req.Phone)
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
