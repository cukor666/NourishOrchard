package adminsvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/admindao"
	"server/models"
)

// ListWithPage 查询管理员列表，分页查询
func ListWithPage(p simpletool.Page, admin models.Admin) ([]models.Admin, int64, error) {
	result, total, err := admindao.ListWithPage(p, admin)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}
