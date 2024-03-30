package empsvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/empdao"
	"server/models"
)

// ListWithPage 查询员工信息，带有分页
func ListWithPage(p simpletool.Page, employee models.Employee) ([]models.Employee, int64, error) {
	result, total, err := empdao.ListWithPage(p, employee)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}
