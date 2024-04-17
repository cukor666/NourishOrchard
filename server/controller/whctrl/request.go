package whctrl

import "server/models"

// insertReq 插入仓库请求结构体
type insertReq struct {
	ID        int64   `json:"id" gorm:"primarykey" form:"id" binding:"omitempty,gt=0"`
	Address   string  `json:"address" form:"address" binding:"required,min=3,max=100"`
	Capacity  float64 `json:"capacity" form:"capacity" binding:"required,gt=700,lt=10000"`
	Remaining float64 `json:"remaining" form:"remaining" binding:"required,gt=0,lt=10000"`
}

// toWarehouse 转换为仓库模型
func (i *insertReq) toWarehouse() models.Warehouse {
	return models.Warehouse{
		ID:        i.ID,
		Address:   i.Address,
		Capacity:  i.Capacity,
		Remaining: i.Remaining,
	}
}

// updateReq 更新仓库请求结构体
type updateReq struct {
	ID        int64   `json:"id" gorm:"primarykey" form:"id" binding:"required,gt=0"`
	Address   string  `json:"address" form:"address" binding:"omitempty,min=3,max=100"`
	Capacity  float64 `json:"capacity" form:"capacity" binding:"omitempty,gt=700,lt=10000"`
	Remaining float64 `json:"remaining" form:"remaining" binding:"omitempty,gt=0,lt=10000"`
}

// toWarehouse 转换为仓库模型
func (u *updateReq) toWarehouse() models.Warehouse {
	return models.Warehouse{
		ID:        u.ID,
		Address:   u.Address,
		Capacity:  u.Capacity,
		Remaining: u.Remaining,
	}
}
