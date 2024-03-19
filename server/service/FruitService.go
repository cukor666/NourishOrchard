package service

import (
	"server/dao"
	"server/models"
)

// List 查询水果列表
func (f FruitService) List() (fruits []models.Fruit, err error) {
	result, err := dao.FruitDao{}.List()
	if err != nil {
		levelLog("查询水果列表失败")
		return nil, err
	}
	return result, nil
}
