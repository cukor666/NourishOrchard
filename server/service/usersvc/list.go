package usersvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/userdao"
	"server/models"
)

// List 访问用户列表，需要管理员或员工权限
func List(p simpletool.Page, user models.User) ([]models.User, int64, error) {
	result, total, err := userdao.ListWithPage(p, user)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}
