package models

type Warehouse struct {
	ID        int64  `json:"id" gorm:"primarykey" form:"id"`
	Address   string `json:"address" form:"address"`
	Capacity  int    `json:"capacity" form:"capacity"`
	Remaining int    `json:"remaining" form:"remaining"`
}

func (w *Warehouse) TableName() string {
	return "warehouse"
}
