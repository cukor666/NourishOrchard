package service

import (
	"errors"
	"regexp"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// Info 根据账号获取管理员信息
func (a AdminService) Info(username string) (admin models.Admin, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Admin{}, errors.New("账号不符合系统规范")
	}
	return dao.AdminDao{}.SelectByUsername(username)
}

// Update 更新管理员信息
func (a AdminService) Update(admin models.Admin) error {
	_, err := dao.AdminDao{}.Update(admin)
	return err
}

// ListWithPage 查询管理员列表，分页查询
func (a AdminService) ListWithPage(p simpletool.Page) ([]models.Admin, int64, error) {
	result, total, err := dao.AdminDao{}.ListWithPage(p)
	if err != nil {
		levelLog("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}
