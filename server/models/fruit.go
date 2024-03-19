package models

type Fruit struct {
	ID         int    `json:"id" gorm:"primarykey" form:"id"`
	Name       string `json:"name" form:"name"`
	Water      int    `json:"water" form:"water"`
	Sugar      int    `json:"sugar" form:"sugar"`
	ShelfLife  int    `json:"shelfLife" form:"shelfLife"`
	Origin     string `json:"origin" form:"origin"`
	SupplierId int    `json:"supplierId" form:"supplierId"`
}

func (f Fruit) TableName() string {
	return "fruits"
}
