package supsvc

import (
	"fmt"
	"server/common/levellog"
	"server/dao/supdao"
	"server/models"
)

func Update(supplier models.Supplier) error {
	err := supdao.Update(supplier)
	if err != nil {
		levellog.Service(fmt.Sprintf("删除失败，sup = %v", supplier))
		return err
	}
	return nil
}
