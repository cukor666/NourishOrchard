package moudels

import "gorm.io/gorm"

type Fruit struct {
	gorm.Model
	FruitName string  `json:"fruitname" gorm:"size:20;not null"`
	Water     uint8   `json:"water"`
	Sugar     uint8   `json:"sugar"`
	Inventory int     `json:"inventory" gorm:"not null"`
	Price     float64 `json:"price" gorm:"not null"`
	SuppherID uint    `json:"suppherid"`
}

// TableName 设置表名
func (f Fruit) TableName() string {
	return "fruits"
}
