package fruitdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
	"server/response"
)

// Detail 水果详情
func Detail(id int) (res response.FruitDetailResponse, err error) {
	var fm models.Fruit
	var sm models.Supplier
	db := dao.GetMySQL()
	err = db.Table(fmt.Sprintf("%s f", fm.TableName())).
		Select("f.id fId", "f.`name` fName", "f.water fWater", "f.sugar fSugar",
			"f.shelf_life fShelfLife", "f.origin fOrigin", "s.id sId", "s.`name` sName",
			"s.address sAddress", "s.contact_person sContactPerson", "s.phone sPhone", "s.reputation sReputation").
		Joins(fmt.Sprintf("JOIN %s s ON f.supplier_id = s.id", sm.TableName())).
		Where("f.id = ?", id).Take(&res).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("查询失败res = %v", res))
	}
	return
}
