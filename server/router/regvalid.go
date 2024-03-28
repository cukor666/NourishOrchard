package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"server/common/levellog"
	"server/router/valid"
)

func RegisterValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			if err := v.RegisterValidation("username", valid.UsernameValid); err != nil {
				levellog.Valid(fmt.Sprintf("注册失败，err: %s", err.Error()))
				panic(err)
			}
		}
	}
}
