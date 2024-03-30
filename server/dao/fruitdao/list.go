package fruitdao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// List 水果列表
func List(p simpletool.Page, fruit models.Fruit) (result []models.Fruit, total int64, err error) {
	db := dao.GetMySQL()
	id, name, water, sugar, life, origin, supplierId := fruit.SetZero()
	fruit.ID, fruit.Water, fruit.Sugar, fruit.ShelfLife, fruit.SupplierId = id, water, sugar, life, supplierId
	tx := db.Model(&fruit).Where(&fruit).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("origin LIKE ?", fmt.Sprintf("%%%s%%", origin)).
		Count(&total)
	levellog.Dao(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levellog.Dao("查询水果列表失败")
		return nil, 0, err
	}
	return result, total, nil
}
