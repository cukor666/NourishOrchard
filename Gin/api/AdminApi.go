package api

import (
	"Gin/dao"
	"Gin/models"
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
	users, result := adminDao.SelectAllAdminByLimit(pageSize, currentPage) // 新版，分页查询
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
	// id, _ := strconv.Atoi(c.Query("ID"))
	// promise, _ := strconv.Atoi(c.Query("promise")) // 删除操作比较危险，只有权限大于普通用户的才可以删除
	type DelteAdminType struct {
		ID      int `json:"ID"`
		Promise int `json:"promise"`
	}

	var temp DelteAdminType
	c.BindJSON(&temp)
	log.Printf("id = %v", temp.ID)
	log.Printf("promise = %v", temp.Promise)
	if temp.Promise < 3 {
		response.Failed("删除失败，权限不够", c)
		return
	}
	var adminDao dao.AdminDao // 这个地方待优化
	ok := adminDao.DeleteById(uint(temp.ID))
	if !ok {
		response.Failed("删除失败，参数错误", c)
	} else {
		response.Success("删除成功", temp.ID, c)
	}
}

// 添加管理员，实际上只是修改用户权限
func AddAdmin(c *gin.Context) {
	var users []models.User
	var ids []uint
	c.ShouldBind(&users)
	// log.Printf("%v", users)
	for _, v := range users {
		ids = append(ids, v.ID)
	}
	// log.Printf("ids = %v", ids)
	var adminDao dao.AdminDao
	err := adminDao.AddAdmin(ids)
	if err != nil {
		log.Printf("add admin failed, err: %v\n", err)
		response.Failed("添加管理员失败", c)
		return
	}
	response.Success("添加管理员成功", ids, c)
}
