package controller

import (
	"github.com/gin-gonic/gin"
)

type Acters interface {
	Info(ctx *gin.Context, username string)
	Update(ctx *gin.Context, username string)
}
