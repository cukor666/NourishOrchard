package request

import "server/models"

type FruitInsertReq struct {
	Name       string `json:"name" binding:"required,max=20"`
	Water      int    `json:"water" binding:"required,gt=0,lt=100"`
	Sugar      int    `json:"sugar" binding:"required,gt=0,lt=100"`
	ShelfLife  int    `json:"shelfLife" binding:"required,gt=0,lt=90"`
	Origin     string `json:"origin" binding:"required,min=3,max=20"`
	SupplierId int    `json:"supplierId" binding:"required,gt=0,lt=20"`
}

func (f *FruitInsertReq) ToFruit() models.Fruit {
	return models.Fruit{
		ID:         0,
		Name:       f.Name,
		Water:      f.Water,
		Sugar:      f.Sugar,
		ShelfLife:  f.ShelfLife,
		Origin:     f.Origin,
		SupplierId: f.SupplierId,
	}
}
