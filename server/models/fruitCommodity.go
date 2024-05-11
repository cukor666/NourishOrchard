package models

type FruitCommodity struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Imgs  string  `json:"imgs"`
	Desc  string  `json:"desc"`
}

func (f *FruitCommodity) TableName() string {
	return "fruits_commodity"
}
