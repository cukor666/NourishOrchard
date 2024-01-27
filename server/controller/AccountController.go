package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"server/models"
	"server/response"
)

func (a AccountController) Insert(c *gin.Context) {
	var account models.Account
	err := c.ShouldBind(&account)
	if err != nil {
		log.Println("绑定失败")
		response.Failed(c, "绑定失败")
		return
	}
	//token := accountService.Insert(account)
	token := ""
	if token == "" {
		response.Failed(c, "插入数据失败")
	}
}
