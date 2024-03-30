package adminsvc

import (
	"server/dao/admindao"
	"server/models"
)

// Update 更新管理员信息
func Update(admin models.Admin) error {
	_, err := admindao.Update(admin)
	return err
}
