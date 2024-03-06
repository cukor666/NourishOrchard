package dao

import "server/models"

type AdminDao struct{}

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
