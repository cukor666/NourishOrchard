package orderctrl

import "server/models"

type insertReq struct {
	Title        string  `json:"title" binding:"required"`
	Status       int     `json:"status" binding:"omitempty,oneof=1"`
	CommodityID  int64   `json:"commodityId" binding:"required"`
	Quantity     float64 `json:"quantity" binding:"required"`
	ReceiverName string  `json:"receiverName" binding:"required"`
	Address      string  `json:"address" binding:"required"`
	Remark       string  `json:"remark" binding:"omitempty"`
}

func (r insertReq) toOrder() models.Order {
	return models.Order{
		Title:        r.Title,
		Status:       r.Status,
		CommodityID:  r.CommodityID,
		Quantity:     r.Quantity,
		ReceiverName: r.ReceiverName,
		Address:      r.Address,
		Remark:       r.Remark,
	}
}
