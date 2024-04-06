package supdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

func Delete(id int64) (sup models.Supplier, err error) {
	db := dao.GetMySQL()
	err = db.Model(&sup).Where("id = ?", id).Delete(&sup).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("从%s表中删除id=%d的数据失败", sup.TableName(), id))
	}
	return
}
