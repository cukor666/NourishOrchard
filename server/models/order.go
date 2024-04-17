package models

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID           int64          `json:"id" form:"id" gorm:"primarykey"`
	Title        string         `json:"title" form:"title"`
	Status       int            `json:"status" form:"status"`
	CommodityID  int64          `json:"commodityId" form:"commodityId"`
	Quantity     float64        `json:"quantity" form:"quantity"`
	BuyerID      int64          `json:"buyerId" form:"buyerId"`
	AdminId      int64          `json:"adminId" form:"adminId"`
	WarehouseID  int64          `json:"warehouseId" form:"warehouseId"`
	ReceiverName string         `json:"receiverName" form:"receiverName"`
	Address      string         `json:"address" form:"address"`
	Remark       string         `json:"remark" form:"remark"`
	CreatedAt    time.Time      `json:"createdAt" form:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt" form:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" form:"deletedAt" gorm:"index"`
}

func (o *Order) TableName() string {
	return "orders"
}

func (o *Order) SetZero() (title, ReceiverName, address, remark string) {
	title, ReceiverName, address, remark = o.Title, o.ReceiverName, o.Address, o.Remark
	o.Title, o.ReceiverName, o.Address, o.Remark = "", "", "", ""
	return
}
