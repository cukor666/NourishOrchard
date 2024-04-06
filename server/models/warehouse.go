package models

type Warehouse struct {
	ID       int64  `json:"id" gorm:"primarykey" form:"id"`
	Address  string `json:"address" form:"address"`
	Capacity int    `json:"capacity" form:"capacity"`
	Status   int    `json:"status" form:"status"`
}

func (w *Warehouse) TableName() string {
	return "warehouse"
}
