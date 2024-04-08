package models

import (
	"gorm.io/gorm"
	"time"
)

type Purchase struct {
	ID         int64           `json:"id" form:"id"`
	Status     int             `json:"status" form:"status"`
	EmployeeId int64           `json:"employeeId" form:"employeeId"`
	FruitId    int64           `json:"fruitId" form:"fruitId"`
	Quantity   int             `json:"quantity" form:"quantity"`
	CreatedAt  time.Time       `json:"createdAt" form:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt" form:"updatedAt"`
	DeletedAt  *gorm.DeletedAt `json:"deletedAt" gorm:"index" form:"deletedAt"`
}

func (p Purchase) TableName() string {
	return "purchase"
}
