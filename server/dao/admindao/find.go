package admindao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// ExistById 根据id查询管理员是否存在
func ExistById(id int64) (bool, error) {
	db := dao.GetMySQL()
	var count int64
	if err := db.Model(&models.Admin{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// SelectByUsername 根据账号查找管理员
func SelectByUsername(username string) (admin models.Admin, err error) {
	db := dao.GetMySQL()
	err = db.Model(&admin).Where("username = ?", username).Take(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// SelectByUsernameAndEmail 根据账号和邮箱查找管理员
func SelectByUsernameAndEmail(username, email string) (admin models.Admin, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.Admin{}).Where("username = ? AND email = ?", username, email).Take(&admin).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("查询管理员失败，admin = %v", admin))
		return models.Admin{}, err
	}
	return admin, nil
}
