package supdao

import (
	"server/dao"
	"server/models"
)

func Update(supplier models.Supplier) error {
	db := dao.GetMySQL()
	return db.Model(&supplier).Where("id = ?", supplier.ID).Omit("id").Updates(&supplier).Error
}
