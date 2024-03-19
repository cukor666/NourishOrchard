package dao

import (
	"server/models"
)

// List 水果列表
func (f FruitDao) List() (result []models.Fruit, err error) {
	err = mysqlDB.Model(&models.Fruit{}).Find(&result).Error
	if err != nil {
		levelLog("查询水果列表失败")
		return nil, err
	}
	return result, nil
}
