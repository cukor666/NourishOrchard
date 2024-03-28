package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"server/router/valid"
)

func RegisterValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			_ = v.RegisterValidation("username", valid.UsernameValid)
			_ = v.RegisterValidation("promise", valid.PromiseValid)
			_ = v.RegisterValidation("gender", valid.GenderValid)
			_ = v.RegisterValidation("phone", valid.PhoneValid)
			_ = v.RegisterValidation("birthday", valid.BirthdayValid)
		}
	}
}
