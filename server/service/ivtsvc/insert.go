package ivtsvc

import (
	"errors"
	"fmt"
	"server/common/levellog"
	"server/dao/ivtdao"
	"server/models"
)

func Insert(inventory models.Inventory) error {
	if ok := has(inventory); !ok {
		levellog.Service("缺少必要添加")
		return errors.New("缺少必要条件")
	}
	if err := ivtdao.Insert(inventory); err != nil {
		levellog.Service(fmt.Sprintf("添加失败，err: %s", err.Error()))
		return err
	}
	return nil
}
