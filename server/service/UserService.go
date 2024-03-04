package service

import (
	"errors"
	"regexp"
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
