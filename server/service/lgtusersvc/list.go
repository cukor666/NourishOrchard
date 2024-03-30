package lgtusersvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/lgtuserdao"
	"server/models"
)

// LogoutList 查询注销用户列表，需要管理员或员工权限
func LogoutList(p simpletool.Page, logoutUser models.LogoutUser) ([]models.LogoutUser, int64, error) {
	result, total, err := lgtuserdao.LogoutListWithPage(p, logoutUser)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}
