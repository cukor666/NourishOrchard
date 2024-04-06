package fruitsvc

import (
	"server/dao/fruitdao"
	"server/models"
)

// Update 更新水果信息
func Update(fruit models.Fruit) error {
	return fruitdao.Update(fruit)
}
