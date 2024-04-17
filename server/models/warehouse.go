package models

type Warehouse struct {
	ID        int64   `json:"id" gorm:"primarykey" form:"id"`
	Address   string  `json:"address" form:"address"`
	Capacity  float64 `json:"capacity" form:"capacity"`
	Remaining float64 `json:"remaining" form:"remaining"`
}

func (w *Warehouse) TableName() string {
	return "warehouse"
}
