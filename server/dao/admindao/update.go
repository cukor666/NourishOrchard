package admindao

import (
	"server/dao"
	"server/models"
)

// Update 更新管理员信息
func Update(admin models.Admin) (models.Admin, error) {
	db := dao.GetMySQL()
	err := db.Model(&admin).Omit("id", "username").Where("username = ?", admin.Username).Updates(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}
