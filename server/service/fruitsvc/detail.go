package fruitsvc

import (
	"server/common/levellog"
	"server/dao/fruitdao"
	"server/response"
)

// Detail 根据id查询水果详细信息
func Detail(id int) (res response.FruitDetailResponse, err error) {
	res, err = fruitdao.Detail(id)
	if err != nil {
		levellog.Service("查询错误")
	}
	return
}
