package fruitsvc

import (
	"server/dao/fruitdao"
	"server/models"
)

// Insert 添加水果信息
func Insert(fruit models.Fruit) error {
	return fruitdao.Insert(fruit)
}
