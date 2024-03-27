package dao

import (
	"fmt"
	"server/common/simpletool"
	"server/models"
	"server/response"
)

// List 水果列表
func (f FruitDao) List(p simpletool.Page, fruit models.Fruit) (result []models.Fruit, total int64, err error) {
	id, name, water, sugar, life, origin, supplierId := fruit.SetZero()
	fruit.ID, fruit.Water, fruit.Sugar, fruit.ShelfLife, fruit.SupplierId = id, water, sugar, life, supplierId
	tx := mysqlDB.Model(&fruit).Where(&fruit).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("origin LIKE ?", fmt.Sprintf("%%%s%%", origin)).
		Count(&total)
	levelLog(fmt.Sprintf("total = %d", total))
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&result).Error
	if err != nil {
		levelLog("查询水果列表失败")
		return nil, 0, err
	}
	return result, total, nil
}

// Detail 水果详情
func (f FruitDao) Detail(id int) (res response.FruitDetailResponse, err error) {
	var fm models.Fruit
	var sm models.Supplier
	err = mysqlDB.Table(fmt.Sprintf("%s f", fm.TableName())).
		Select("f.id fId", "f.`name` fName", "f.water fWater", "f.sugar fSugar",
			"f.shelf_life fShelfLife", "f.origin fOrigin", "s.id sId", "s.`name` sName",
			"s.address sAddress", "s.contact_person sContactPerson", "s.phone sPhone", "s.reputation sReputation").
		Joins(fmt.Sprintf("JOIN %s s ON f.supplier_id = s.id", sm.TableName())).
		Where("f.id = ?", id).Take(&res).Error
	if err != nil {
		levelLog(fmt.Sprintf("查询失败res = %v", res))
	}
	return
}

// Insert 插入水果信息
func (f FruitDao) Insert(fruit models.Fruit) error {
	err := mysqlDB.Model(&models.Fruit{}).Create(&fruit).Error
	if err != nil {
		levelLog("插入水果信息失败")
		return err
	}
	return nil
}

// Delete 根据Id删除水果信息
func (f FruitDao) Delete(id int) (fruit models.Fruit, err error) {
	tx := mysqlDB.Model(&fruit).Where("id = ?", id).Take(&fruit)
	err = tx.Error
	if err != nil {
		levelLog(fmt.Sprintf("%s表中无%d的数据", fruit.TableName(), id))
		return models.Fruit{}, err
	}
	err = tx.Model(&fruit).Where("id = ?", id).Delete(&models.Fruit{}).Error
	if err != nil {
		levelLog(fmt.Sprintf("无法从%s表中删除%v的数据", fruit.TableName(), fruit))
		return models.Fruit{}, err
	}
	return fruit, nil
}

// Update 更新水果信息
func (f FruitDao) Update(fruit models.Fruit) error {
	err := mysqlDB.Model(&models.Fruit{}).Where("id = ?", fruit.ID).Omit("id").Updates(&fruit).Error
	if err != nil {
		levelLog("更新失败")
		return err
	}
	return nil
}
