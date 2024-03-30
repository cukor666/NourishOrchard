package fruitdao

import (
	"fmt"
	"server/common/levellog"
	"server/dao"
	"server/models"
)

// Delete 根据Id删除水果信息
func Delete(id int) (fruit models.Fruit, err error) {
	db := dao.GetMySQL()
	tx := db.Model(&fruit).Where("id = ?", id).Take(&fruit)
	err = tx.Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("%s表中无%d的数据", fruit.TableName(), id))
		return models.Fruit{}, err
	}
	err = tx.Model(&fruit).Where("id = ?", id).Delete(&models.Fruit{}).Error
	if err != nil {
		levellog.Dao(fmt.Sprintf("无法从%s表中删除%v的数据", fruit.TableName(), fruit))
		return models.Fruit{}, err
	}
	return fruit, nil
}
