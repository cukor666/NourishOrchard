package dao

import (
	"Gin/app"
	"Gin/models"
	"log"
)

type FruitDao struct{}

// Insert 添加新水果
func (fd FruitDao) Insert(fruit models.Fruit) (newId uint, affect int64) {
	d := app.MySQLDB().Create(&fruit)
	if d.Error != nil {
		log.Printf("insert fruit failed, err: %v\n", d.Error)
		return 0, 0
	}
	return fruit.ID, d.RowsAffected
}

// SelectAll 查看所有水果
func (fd FruitDao) SelectAll() (fruits []models.Fruit, ok bool) {
	err := app.MySQLDB().Omit("DeletedAt").Find(&fruits).Error
	if err != nil {
		log.Printf("select fruit all failed, err: %v\n", err)
		return nil, false
	}
	return fruits, true
}
