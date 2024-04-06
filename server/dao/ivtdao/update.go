package ivtdao

import (
	"errors"
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Update 更新库存信息
func Update(inventory models.Inventory) error {
	tx := dao.GetMySQL().Begin()
	cnts := [3]int64{}

	if inventory.CommodityId != 0 {
		err := tx.Model(&models.Fruit{}).Where("id = ?", inventory.CommodityId).Count(&cnts[0]).Error
		if err != nil || cnts[0] == 0 {
			w := fmt.Sprintf("查询水果id=%d失败", inventory.CommodityId)
			levellog.Dao(w)
			tx.Rollback()
			return errors.New(w)
		}
	}

	if inventory.EmployeeId != 0 {
		err := tx.Model(&models.Employee{}).Where("id = ?", inventory.EmployeeId).Count(&cnts[1]).Error
		if err != nil || cnts[1] == 0 {
			w := fmt.Sprintf("查询员工id=%d失败", inventory.EmployeeId)
			levellog.Dao(w)
			tx.Rollback()
			return errors.New(w)
		}
	}

	if inventory.WarehouseId != 0 {
		err := tx.Model(&models.Warehouse{}).Where("id = ?", inventory.WarehouseId).Count(&cnts[2]).Error
		if err != nil || cnts[2] == 0 {
			w := fmt.Sprintf("查询仓库id=%d失败", inventory.WarehouseId)
			levellog.Dao(w)
			tx.Rollback()
			return errors.New(w)
		}
	}

	err := tx.Model(&models.Inventory{}).Where("id = ?", inventory.ID).Omit("id").Updates(&inventory).Error

	if err != nil {
		levellog.Dao("更新失败")
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
