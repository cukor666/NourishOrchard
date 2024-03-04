package models

type Fruit struct {
	ID         int    `json:"id" gorm:"primarykey"`
	Name       string `json:"name"`
	Water      int    `json:"water"`
	Sugar      int    `json:"sugar"`
	ShelfLife  int    `json:"shelfLife"`
	Origin     string `json:"origin"`
	SupplierId int    `json:"supplierId"`
}

func (f Fruit) TableName() string {
	return "fruit"
}
