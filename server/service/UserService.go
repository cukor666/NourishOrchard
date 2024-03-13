package service

import (
	"errors"
	"regexp"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// Info 根据用户名获取用户信息
func (u UserService) Info(username string) (user models.User, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.User{}, errors.New("用户名不符合系统规范")
	}
	return dao.UserDao{}.SelectByUsername(username)
}

// Update 更新用户信息
func (u UserService) Update(user models.User) error {
	_, err := dao.UserDao{}.Update(user)
	return err
}

// List 访问用户列表，需要管理员或员工权限
func (u UserService) List(p simpletool.Page, user models.User) ([]models.User, int64, error) {
	result, total, err := dao.UserDao{}.ListWithPage(p, user)
	if err != nil {
		levelLog("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// DeleteUser 根据账号删除用户信息
func (u UserService) DeleteUser(username string) (models.User, error) {
	user, err := dao.UserDao{}.DeleteByUsername(username)
	if err != nil {
		levelLog("用户删除失败")
		return models.User{}, err
	}
	return user, nil
}

// LogoutList 查询注销用户列表，需要管理员或员工权限
func (u UserService) LogoutList(p simpletool.Page, logoutUser models.LogoutUser) ([]models.LogoutUser, int64, error) {
	result, total, err := dao.UserDao{}.LogoutListWithPage(p, logoutUser)
	if err != nil {
		levelLog("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// RecoverUser 恢复用户信息
func (u UserService) RecoverUser(username string) (models.User, error) {
	user, err := dao.UserDao{}.RecoverUser(username)
	if err != nil {
		levelLog("无法恢复用户信息")
		return models.User{}, err
	}
	return user, nil
}
