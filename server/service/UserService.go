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
	return userDao.SelectByUsername(username)
}

// Update 更新用户信息
func (u UserService) Update(user models.User) error {
	_, err := dao.UserDao{}.Update(user)
	return err
}

// List 访问用户列表，需要管理员或员工权限
func (u UserService) List(p simpletool.Page) ([]models.User, error) {
	//result, err := dao.UserDao{}.List()	// 不带分页
	result, err := dao.UserDao{}.ListWithPage(p)
	if err != nil {
		levelLog("查询失败")
		return nil, err
	}
	return result, nil
}
