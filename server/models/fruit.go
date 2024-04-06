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

func (f *Fruit) TableName() string {
	return "fruits"
}

// SetZero 返回对应字段原有的值，并设置字段为零值
func (f *Fruit) SetZero() (id int, name string, water, sugar, shelfLife int, origin string, supplierId int) {
	id, name, water, sugar, shelfLife, origin, supplierId = f.ID, f.Name, f.Water, f.Sugar, f.ShelfLife, f.Origin, f.SupplierId
	f.ID, f.Name, f.Water, f.Sugar, f.ShelfLife, f.Origin, f.SupplierId = 0, "", 0, 0, 0, "", 0
	return
}
