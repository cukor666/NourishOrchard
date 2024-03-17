package service

import (
	"errors"
	"fmt"
	"regexp"
	"server/dao"
	"server/request"
	"server/utils"
)

// IsExists 判断账号在数据库中是否存在
func (a AccountService) IsExists(username string, promise int) bool {
	// 如果格式不对的话直接返回false
	prefix := utils.GetUsernamePrefix()
	compile := regexp.MustCompile("^" + prefix + "[0-9]+$")
	ok := compile.MatchString(username)
	if !ok {
		return false
	}
	// 开始查询数据库
	cnt := dao.AccountDao{}.GetCountByUsername(username, promise)
	return cnt != 0
}

// Exit 账号退出
func (a AccountService) Exit(username string) error {
	err := dao.AccountDao{}.Exit(username)
	if err != nil {
		levelLog("退出失败")
		return err
	}
	return nil
}

// ChangePassword 修改密码
func (a AccountService) ChangePassword(username, oldPassword, newPassword string) error {
	// 从数据库中获取原密码
	password, err := dao.AccountDao{}.GetPassword(username)
	if err != nil {
		levelLog("获取密码失败")
		return err
	}
	// 将前端传递的密码和数据库中的密码进行比较
	if !utils.PwdOK(password, oldPassword) {
		levelLog("前端传递密码与数据库中原密码不一致，拒绝修改密码请求")
		return errors.New("前端传递密码与数据库中密码不一致")
	}
	// 对新密码进行加密
	newPwd, err := utils.GetPwd(newPassword)
	if err != nil {
		levelLog(fmt.Sprintf("对密码加密失败, newPassword = %s", newPassword))
		return err
	}
	err = dao.AccountDao{}.ChangePassword(username, string(newPwd))
	if err != nil {
		levelLog("修改密码失败")
		return err
	}
	return nil
}

// ForgetPassword 用户忘记密码
func (a AccountService) ForgetPassword(req request.ForgetPwdReq) (err error) {
	// 校验数据库是否有该用户并且校验绑定的号码是否正确
	_, err = dao.UserDao{}.SelectByUsernameAndPhone(req.Username, req.Phone)
	if err != nil {
		levelLog("用户绑定电话错误")
		return err
	}
	// 对新密码加密
	pwd, err := utils.GetPwd(req.Password)
	if err != nil {
		levelLog("新密码加密失败")
		return err
	}
	// 更新新密码
	err = dao.AccountDao{}.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levelLog("修改密码错误")
		return err
	}
	return nil
}
