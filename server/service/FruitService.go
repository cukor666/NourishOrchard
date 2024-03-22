package service

import (
	"server/common/simpletool"
	"server/dao"
	"server/models"
	"server/response"
)

// List 查询水果列表
func (f FruitService) List(p simpletool.Page, fruit models.Fruit) (fruits []models.Fruit, total int64, err error) {
	fruits, total, err = dao.FruitDao{}.List(p, fruit)
	if err != nil {
		levelLog("查询水果列表失败")
		return nil, 0, err
	}
	return fruits, total, nil
}

// Detail 根据id查询水果详细信息
func (f FruitService) Detail(id int) (res response.FruitDetailResponse, err error) {
	res, err = dao.FruitDao{}.Detail(id)
	if err != nil {
		levelLog("查询错误")
	}
	return
}

// Insert 添加水果信息
func (f FruitService) Insert(fruit models.Fruit) error {
	return dao.FruitDao{}.Insert(fruit)
}

// Delete 删除水果信息
func (f FruitService) Delete(id int) (fruit models.Fruit, err error) {
	return dao.FruitDao{}.Delete(id)
}

// Update 更新水果信息
func (f FruitService) Update(fruit models.Fruit) error {
	return dao.FruitDao{}.Update(fruit)
}
