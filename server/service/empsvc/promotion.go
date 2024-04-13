package empsvc

import (
	"errors"
	"fmt"
	"server/dao/actdao"
	"server/dao/empdao"
	"server/models"
	mc "server/models/code"
)

// Promotion 员工晋升管理员
func Promotion(username, name, email, mark string) error {
	if cnt := actdao.GetCountByUsername(username, mc.EMPLOYEE); cnt == 0 {
		return errors.New(fmt.Sprintf("%s表中无数据username = %s", (&models.Account{}).TableName(), username))
	}
	return empdao.Promotion(username, name, email, mark)
}
