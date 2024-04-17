package ivtctrl

import "server/models"

type insertReq struct {
	ID          int64   `json:"id" form:"id" gorm:"primarykey" binding:"omitempty,gt=0"` // 库存ID
	CommodityId int64   `json:"commodityId" form:"commodityId" binding:"required,gt=0"`  // 水果ID
	Quantity    float64 `json:"quantity" form:"quantity" binding:"required,gt=0"`        // 水果数量，箱
	EmployeeId  int64   `json:"employeeId" form:"employeeId" binding:"required,gt=0"`    // 操作的员工ID
	WarehouseId int64   `json:"warehouseId" form:"warehouseId" binding:"required,gt=0"`  // 仓库ID
}

func (i *insertReq) toInventory() models.Inventory {
	return models.Inventory{
		ID:          i.ID,
		CommodityId: i.CommodityId,
		Quantity:    i.Quantity,
		EmployeeId:  i.EmployeeId,
		WarehouseId: i.WarehouseId,
	}
}

type updateReq struct {
	ID          int64   `json:"id" binding:"required,gt=0"`            // 库存ID
	CommodityId int64   `json:"commodityId"  binding:"omitempty,gt=0"` // 水果ID
	Quantity    float64 `json:"quantity" binding:"omitempty,gt=0"`     // 水果数量，箱
	EmployeeId  int64   `json:"employeeId"  binding:"omitempty,gt=0"`  // 操作的员工ID
	WarehouseId int64   `json:"warehouseId" binding:"omitempty,gt=0"`  // 仓库ID
}

func (i *updateReq) toInventory() models.Inventory {
	return models.Inventory{
		ID:          i.ID,
		CommodityId: i.CommodityId,
		Quantity:    i.Quantity,
		EmployeeId:  i.EmployeeId,
		WarehouseId: i.WarehouseId,
	}
}
