package dao

import (
	"Gin/moudels"
	"log"
	"time"
)

type FruitDao struct{}

//func init() {
//	db.AutoMigrate(&moudels.Fruit{})
//}

// Insert 添加新水果
func (fd FruitDao) Insert(fruit moudels.Fruit) (newId uint, affect int64) {
	fruit.CreatedAt = time.Now()
	fruit.UpdatedAt = time.Now()
	d := db.Create(&fruit)
	if d.Error != nil {
		log.Printf("insert fruit failed, err: %v\n", d.Error)
		return 0, 0
	}
	return fruit.ID, d.RowsAffected
}

// SelectAll 查看所有水果
func (fd FruitDao) SelectAll() (fruits []moudels.Fruit, ok bool) {
	ok = true
	err := db.Omit("DeletedAt").
		Find(&fruits).Error
	if err != nil {
		log.Printf("select fruit all failed, err: %v\n", err)
		ok = false
		return
	}
	return
}
