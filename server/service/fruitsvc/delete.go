package fruitsvc

import (
	"server/dao/fruitdao"
	"server/models"
)

// Delete 删除水果信息
func Delete(id int) (fruit models.Fruit, err error) {
	return fruitdao.Delete(id)
}
