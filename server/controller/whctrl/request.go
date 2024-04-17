package whctrl

import "server/models"

type insertReq struct {
	ID        int64  `json:"id" gorm:"primarykey" form:"id" binding:"omitempty,gt=0"`
	Address   string `json:"address" form:"address" binding:"required,min=3,max=100"`
	Capacity  int    `json:"capacity" form:"capacity" binding:"required,gt=700,lt=10000"`
	Remaining int    `json:"remaining" form:"remaining" binding:"required,gt=0,lt=10000"`
}

func (i *insertReq) toWarehouse() models.Warehouse {
	return models.Warehouse{
		ID:        i.ID,
		Address:   i.Address,
		Capacity:  i.Capacity,
		Remaining: i.Remaining,
	}
}

type updateReq struct {
	ID        int64  `json:"id" gorm:"primarykey" form:"id" binding:"required,gt=0"`
	Address   string `json:"address" form:"address" binding:"omitempty,min=3,max=100"`
	Capacity  int    `json:"capacity" form:"capacity" binding:"omitempty,gt=700,lt=10000"`
	Remaining int    `json:"remaining" form:"remaining" binding:"omitempty,gt=0,lt=10000"`
}

func (u *updateReq) toWarehouse() models.Warehouse {
	return models.Warehouse{
		ID:        u.ID,
		Address:   u.Address,
		Capacity:  u.Capacity,
		Remaining: u.Remaining,
	}
}
