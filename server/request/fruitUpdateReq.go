package request

import "server/models"

type FruitUpdateReq struct {
	ID         int    `json:"id" binding:"required,gt=0"`
	Name       string `json:"name,omitempty"`
	Water      int    `json:"water,omitempty"`
	Sugar      int    `json:"sugar,omitempty"`
	ShelfLife  int    `json:"shelfLife,omitempty"`
	Origin     string `json:"origin,omitempty"`
	SupplierId int    `json:"supplierId,omitempty"`
}

func (f *FruitUpdateReq) ToFruit() models.Fruit {
	return models.Fruit{
		ID:         f.ID,
		Name:       f.Name,
		Water:      f.Water,
		Sugar:      f.Sugar,
		ShelfLife:  f.ShelfLife,
		Origin:     f.Origin,
		SupplierId: f.SupplierId,
	}
}
