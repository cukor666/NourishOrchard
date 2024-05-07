package fruitcmdydao

import (
	"server/dao"
	"server/dto/fruitdto"
)

func List() (results []fruitdto.FruitCommodity, total int64, err error) {
	db := dao.GetMySQL()
	/*
		sql:
			SELECT fc.id id, f.`name`, fc.price price, fc.imgs imgs, fc.`desc` `desc`, f.origin
			FROM `fruits_commodity` fc
			JOIN fruits f ON f.id = fc.id
	*/
	sql := "SELECT fc.id id, f.`name`, fc.price price, fc.imgs imgs, fc.`desc` `desc`, f.origin" +
		" FROM `fruits_commodity` fc JOIN fruits f ON f.id = fc.id"
	tx := db.Raw(sql).Scan(&results)
	total = tx.RowsAffected
	err = tx.Error
	return
}
