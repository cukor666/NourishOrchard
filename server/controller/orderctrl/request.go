package orderctrl

import "server/models"

type insertReq struct {
	Title         string  `json:"title" binding:"required"`
	Status        int     `json:"status" binding:"omitempty,oneof=1"`
	CommodityID   int64   `json:"commodityId" binding:"required"`
	Quantity      float64 `json:"quantity" binding:"required,gt=0"`
	ReceiverName  string  `json:"receiverName" binding:"required"`
	ReceiverPhone string  `json:"receiverPhone" binding:"required"`
	Address       string  `json:"address" binding:"required"`
	Remark        string  `json:"remark" binding:"omitempty"`
}

func (r insertReq) toOrder() models.Order {
	return models.Order{
		Title:         r.Title,
		Status:        r.Status,
		CommodityID:   r.CommodityID,
		Quantity:      r.Quantity,
		ReceiverName:  r.ReceiverName,
		ReceiverPhone: r.ReceiverPhone,
		Address:       r.Address,
		Remark:        r.Remark,
	}
}

type updateReq struct {
	ID           int64   `json:"id" binding:"required"`
	Title        string  `json:"title" binding:"omitempty"`
	Status       int     `json:"status" binding:"omitempty,oneof=1 2 3 4"`
	CommodityID  int64   `json:"commodityId" binding:"required,gt=0"`
	WarehouseID  int64   `json:"warehouseId" binding:"required,gt=0"`
	Quantity     float64 `json:"quantity" binding:"omitempty,gt=0"`
	ReceiverName string  `json:"receiverName" binding:"omitempty"`
	Address      string  `json:"address" binding:"omitempty"`
	Remark       string  `json:"remark" binding:"omitempty"`
}

func (r updateReq) toOrder() models.Order {
	return models.Order{
		ID:           r.ID,
		Title:        r.Title,
		Status:       r.Status,
		CommodityID:  r.CommodityID,
		WarehouseID:  r.WarehouseID,
		Quantity:     r.Quantity,
		ReceiverName: r.ReceiverName,
		Address:      r.Address,
		Remark:       r.Remark,
	}
}
