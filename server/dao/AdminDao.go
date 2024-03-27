package dao

import (
	"fmt"
	"server/common/simpletool"
	"server/models"
)

func (ad *AdminDao) GetUId(username string) (id uint, err error) {
	var am models.Admin
	err = mysqlDB.Table(am.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}

func (ad *AdminDao) SelectByUsername(username string) (admin models.Admin, err error) {
	err = mysqlDB.Model(&admin).Where("username = ?", username).Take(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// Update 更新管理员信息
func (ad *AdminDao) Update(admin models.Admin) (models.Admin, error) {
	err := mysqlDB.Model(&admin).Omit("id", "username").Where("username = ?", admin.Username).Updates(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// ListWithPage 查询管理员接口
func (ad *AdminDao) ListWithPage(p simpletool.Page) (result []models.Admin, total int64, err error) {
	tx := mysqlDB.Model(&models.Admin{}).Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询管理员列表失败")
		return nil, 0, err
	}
	return result, total, nil
}

// SelectByUsernameAndEmail 根据账号和邮箱查找管理员
func (ad *AdminDao) SelectByUsernameAndEmail(username, email string) (admin models.Admin, err error) {
	err = mysqlDB.Model(&models.Admin{}).Where("username = ? AND email = ?", username, email).Take(&admin).Error
	if err != nil {
		levelLog(fmt.Sprintf("查询管理员失败，admin = %v", admin))
		return models.Admin{}, err
	}
	return admin, nil
}
