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

// 删除用户，根据ID
func DeleteAdmin(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("ID"))
	promise, _ := strconv.Atoi(c.Query("promise")) // 删除操作比较危险，只有权限大于普通用户的才可以删除
	if promise < 3 {
		response.Failed("删除失败，权限不够", c)
		return
	}
	var userDao dao.UserDao // 这个地方待优化
	ok := userDao.DeleteById(uint(id))
	if !ok {
		response.Failed("删除失败，参数错误", c)
	} else {
		response.Success("删除成功", id, c)
	}
}
