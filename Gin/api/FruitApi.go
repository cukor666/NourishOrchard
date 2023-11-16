package api

import (
	"Gin/dao"
	"Gin/moudels"
	"Gin/response"
	"log"

	"github.com/gin-gonic/gin"
)

// 添加水果
func AddFruit(c *gin.Context) {
	var data moudels.Fruit
	c.BindJSON(&data)
	log.Printf("data = %v", data)
	var fruitDao dao.FruitDao
	newId, affect := fruitDao.Insert(data)
	if affect == 0 {
		response.Failed("添加水果失败", c)
	} else {
		response.Success("添加水果成功", newId, c)
	}
}

// 查看所有水果
func SeeAllFruit(c *gin.Context) {
	var fruitDao dao.FruitDao
	fruits, ok := fruitDao.SelectAll()
	if !ok {
		response.Failed("无法获取数据库中的所有水果", c)
	} else {
		response.Success("获取所有水果信息成功", fruits, c)
	}
}
