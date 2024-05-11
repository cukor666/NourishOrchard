package orderctrl

import "server/models"

type insertReq struct {
	Title         string  `json:"title" binding:"required"`           // 订单标题
	Status        int     `json:"status" binding:"omitempty,oneof=1"` // 订单状态，1-待付款，2-待发货，3-待收货，4-已完成
	CommodityID   int64   `json:"commodityId" binding:"required"`     // 商品ID
	Quantity      float64 `json:"quantity" binding:"required,gt=0"`   // 商品数量
	ReceiverName  string  `json:"receiverName" binding:"required"`    // 收货人姓名
	ReceiverPhone string  `json:"receiverPhone" binding:"required"`   // 收货人手机号
	Address       string  `json:"address" binding:"required"`         // 收货地址
	Remark        string  `json:"remark" binding:"omitempty"`         // 备注
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
