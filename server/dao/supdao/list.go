package supdao

import (
	"fmt"
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao"
	"server/models"
)

// List 查询供应商列表，分页查询
func List(p simpletool.Page, supplier models.Supplier) (results []models.Supplier, total int64, err error) {
	id, name, address, person, phone, reputation := supplier.SetZero()
	supplier.ID = id
	supplier.Reputation = reputation
	db := dao.GetMySQL()
	tx := db.Model(&supplier).Where(&supplier).
		Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("address LIKE ?", fmt.Sprintf("%%%s%%", address)).
		Where("contact_person LIKE ?", fmt.Sprintf("%%%s%%", person)).
		Where("phone LIKE ?", fmt.Sprintf("%%%s%%", phone)).
		Count(&total)
	err = tx.Limit(p.Size).Offset((p.Num - 1) * p.Size).Find(&results).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("查询%s表失败", supplier.TableName()))
		return nil, 0, err
	}
	return
}
