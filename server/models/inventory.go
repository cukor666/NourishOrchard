package models

import (
	"gorm.io/gorm"
	"time"
)

type Inventory struct {
	ID          int64           `json:"id" form:"id" gorm:"primarykey"`          // 库存ID
	CommodityId int64           `json:"commodityId" form:"commodityId"`          // 水果ID
	Quantity    int             `json:"quantity" form:"quantity"`                // 水果数量，箱
	EmployeeId  int64           `json:"employeeId" form:"employeeId"`            // 操作的员工ID
	WarehouseId int64           `json:"warehouseId" form:"warehouseId"`          // 仓库ID
	CreatedAt   time.Time       `json:"createdAt"`                               // 创建时间
	UpdatedAt   time.Time       `json:"updatedAt"`                               // 更新时间
	DeletedAt   *gorm.DeletedAt `json:"deletedAt" gorm:"index" form:"deletedAt"` // 删除时间
}

func (i *Inventory) TableName() string {
	return "inventory"
}
