package fruitsvc

import (
	"server/dao/fruitdao"
	"server/dto/fruitdto"
)

// Detail 根据id查询水果详细信息
func Detail(id int) (res fruitdto.Detail, err error) {
	return fruitdao.Detail(id)
}
