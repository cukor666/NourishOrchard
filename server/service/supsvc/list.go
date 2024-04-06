package supsvc

import (
	"server/common/simpletool"
	"server/dao/supdao"
	"server/models"
)

// List 供应商列表，分页查询
func List(p simpletool.Page, supplier models.Supplier) ([]models.Supplier, int64, error) {
	return supdao.List(p, supplier)
}
