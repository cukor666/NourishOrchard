package controller

import (
	"github.com/gin-gonic/gin"
)

type acters interface {
	Info(ctx *gin.Context, username string)
	Update(ctx *gin.Context, username string)
}
