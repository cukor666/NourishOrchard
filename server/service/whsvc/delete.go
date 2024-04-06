package whsvc

import (
	"server/common/levellog"
	"server/dao/whdao"
)

func Delete(id int64) error {
	_, err := whdao.Delete(id)
	if err != nil {
		levellog.Service("删除失败")
		return err
	}
	return nil
}
