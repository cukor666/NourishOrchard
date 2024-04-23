package fruitdao

import (
	"server/dao"
	"server/models"
)

// OriginCount 根据水果的原产地查询水果数量
func OriginCount(origin string) (cnt int64, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.Fruit{}).Where("origin LIKE ?", "%"+origin+"%").Count(&cnt).Error
	return
}
