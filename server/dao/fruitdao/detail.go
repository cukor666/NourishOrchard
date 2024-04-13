package fruitdao

import (
	"fmt"
	"server/dao"
	"server/dto/fruitdto"
	"server/models"
)

// Detail 水果详情
func Detail(id int) (res fruitdto.Detail, err error) {
	err = dao.GetMySQL().Table(fmt.Sprintf("%s f", (&models.Fruit{}).TableName())).
		Select("f.id fId", "f.`name` fName", "f.water fWater", "f.sugar fSugar",
			"f.shelf_life fShelfLife", "f.origin fOrigin", "s.id sId", "s.`name` sName",
			"s.address sAddress", "s.contact_person sContactPerson", "s.phone sPhone", "s.reputation sReputation").
		Joins(fmt.Sprintf("JOIN %s s ON f.supplier_id = s.id", (&models.Supplier{}).TableName())).
		Where("f.id = ?", id).Take(&res).Error
	return
}
