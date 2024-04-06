package ivtsvc

import (
	"server/common/levellog"
	"server/dao/ivtdao"
)

func Delete(id int64) error {
	err := ivtdao.Delete(id)
	if err != nil {
		levellog.Service("删除失败")
		return err
	}
	levellog.Service("删除成功")
	return nil
}
