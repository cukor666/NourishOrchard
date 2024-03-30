package fruitsvc

import (
	"server/common/levellog"
	"server/common/simpletool"
	"server/dao/fruitdao"
	"server/models"
)

// List 查询水果列表
func List(p simpletool.Page, fruit models.Fruit) (fruits []models.Fruit, total int64, err error) {
	fruits, total, err = fruitdao.List(p, fruit)
	if err != nil {
		levellog.Service("查询水果列表失败")
		return nil, 0, err
	}
	return fruits, total, nil
}
