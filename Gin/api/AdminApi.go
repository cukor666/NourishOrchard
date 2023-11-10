package api

import (
	"Gin/dao"
	"Gin/response"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询所有用户信息，分页
func AdminList(c *gin.Context) {
	var adminDao dao.AdminDao
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	log.Printf("pageSize = %v", pageSize)
	log.Printf("currentPage = %v", currentPage)
	users, result := adminDao.SelectAllAdminByLimt(pageSize, currentPage) // 新版，分页查询
	if result.Error != nil {
		log.Printf("select all by limit failed, err: %v\n", result.Error)
		response.Failed("后端查询数据失败", c)
	} else {
		response.Success("查询数据成功", gin.H{
			"users": users,
			"rows":  result.Rows,
		}, c)
	}
}
