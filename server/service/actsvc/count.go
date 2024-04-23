package actsvc

import (
	"fmt"
	"server/common/levellog"
	"server/dao/actdao"
	"server/dto/actdto"
)

func Count() (result []actdto.Count, err error) {
	result, err = actdao.Count()
	if err != nil {
		levellog.Service(fmt.Sprintf("查询失败：%v", err))
		return nil, err
	}
	return result, nil
}
