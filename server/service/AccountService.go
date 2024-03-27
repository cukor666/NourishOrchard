package service

import (
	"errors"
	"fmt"
	"regexp"
	"server/common/levellog"
	"server/models"
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
	cnt := accountDao.GetCountByUsername(username, promise)
	return cnt != 0
}

// Get 获取账号信息
func (a AccountService) Get(username string, promise int) (account models.Account, err error) {
	return accountDao.Get(username, promise)
}

// Exit 账号退出
func (a AccountService) Exit(username string) error {
	err := accountDao.Exit(username)
	if err != nil {
		levellog.Service("退出失败")
		return err
	}
	return nil
}

// ChangePassword 修改密码
func (a AccountService) ChangePassword(username, oldPassword, newPassword string) error {
	// 从数据库中获取原密码
	password, err := accountDao.GetPassword(username)
	if err != nil {
		levellog.Service("获取密码失败")
		return err
	}
	// 将前端传递的密码和数据库中的密码进行比较
	if !utils.PwdOK(password, oldPassword) {
		levellog.Service("前端传递密码与数据库中原密码不一致，拒绝修改密码请求")
		return errors.New("前端传递密码与数据库中密码不一致")
	}
	// 对新密码进行加密
	newPwd, err := utils.GetPwd(newPassword)
	if err != nil {
		levellog.Service(fmt.Sprintf("对密码加密失败, newPassword = %s", newPassword))
		return err
	}
	err = accountDao.ChangePassword(username, string(newPwd))
	if err != nil {
		levellog.Service("修改密码失败")
		return err
	}
	return nil
}
