package empsvc

import (
	"errors"
	"fmt"
	"server/common/levellog"
	"server/dao/actdao"
	"server/dao/empdao"
	"server/models"
	mc "server/models/code"
	"server/request"
)

// Promotion 员工晋升管理员
func Promotion(req request.PromotionRequest) error {
	cnt := actdao.GetCountByUsername(req.Username, mc.EMPLOYEE)
	if cnt == 0 {
		var am models.Account
		str := fmt.Sprintf("%s表中无数据username = %s", am.TableName(), req.Username)
		levellog.Service(str)
		return errors.New(str)
	}
	err := empdao.Promotion(req.Username, req.Name, req.Email, req.Mark)
	if err != nil {
		levellog.Service("晋升管理员失败")
		return err
	}
	return nil
}
