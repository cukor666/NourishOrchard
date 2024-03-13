package models

import (
	"database/sql"
	"time"
)

type Fruit struct {
	ID        uint         `json:"ID" gorm:"primarykey"`
	FruitName string       `json:"fruitname" gorm:"size:20;not null"`
	Water     uint8        `json:"water"`
	Sugar     uint8        `json:"sugar"`
	Inventory int          `json:"inventory" gorm:"not null"`
	Price     float64      `json:"price" gorm:"not null"`
	SuppherID uint         `json:"suppherId"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deletedAt" gorm:"index"`
}

// TableName 设置表名
func (f Fruit) TableName() string {
	return "fruits"
}
