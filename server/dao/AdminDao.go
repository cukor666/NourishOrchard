package dao

import (
	"server/common/simpletool"
	"server/models"
)

func (a AdminDao) GetUId(username string) (id uint, err error) {
	err = mysqlDB.Table(models.Admin{}.TableName()).
		Where("username = ?", username).
		Select("id").Take(&id).Error
	return
}

func (a AdminDao) SelectByUsername(username string) (admin models.Admin, err error) {
	err = mysqlDB.Model(&admin).Where("username = ?", username).Take(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// Update 更新管理员信息
func (a AdminDao) Update(admin models.Admin) (models.Admin, error) {
	err := mysqlDB.Model(&admin).Omit("id", "username").Where("username = ?", admin.Username).Updates(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}
	return admin, nil
}

// ListWithPage 查询管理员接口
func (a AdminDao) ListWithPage(p simpletool.Page) (result []models.Admin, total int64, err error) {
	tx := mysqlDB.Model(&models.Admin{}).Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询管理员列表失败")
		return nil, 0, err
	}
	return result, total, nil
}
