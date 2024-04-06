package supsvc

import (
	"fmt"
	"server/common/levellog"
	"server/dao/supdao"
)

func Delete(id int64) error {
	sup, err := supdao.Delete(id)
	if err != nil {
		levellog.Service("删除失败")
		return err
	}
	levellog.Service(fmt.Sprintf("删除成功sup=%v", sup))
	return nil
}
