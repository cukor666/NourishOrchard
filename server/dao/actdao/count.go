package actdao

import (
	"server/common/levellog"
	"server/dao"
	"server/dto/actdto"
	"server/models"
)

// Count 统计人员数量
func Count() (result []actdto.Count, err error) {
	db := dao.GetMySQL()
	err = db.Model(&models.Account{}).Select("promise, COUNT(*) as cnt").Group("promise").Scan(&result).Error
	if err != nil {
		levellog.Dao("查询失败")
		return nil, err
	}
	return result, nil
}
