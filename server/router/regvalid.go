package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"server/common/levellog"
	"server/router/valid"
)

func RegisterValid(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			levellog.Valid(fmt.Sprintf("校验器异常，err: %v", err))
		}
	}()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("username", valid.UsernameValid); err != nil {
			panic(err)
		}

		if err := v.RegisterValidation("promise", valid.PromiseValid); err != nil {
			panic(err)
		}

		if err := v.RegisterValidation("gender", valid.GenderValid); err != nil {
			panic(err)
		}

		if err := v.RegisterValidation("phone", valid.PhoneValid); err != nil {
			panic(err)
		}

		if err := v.RegisterValidation("birthday", valid.BirthdayValid); err != nil {
			panic(err)
		}

		if err := v.RegisterValidation("email", valid.EmailValid); err != nil {
			panic(err)
		}

		if err := v.RegisterValidation("empPosition", valid.EmpPosition); err != nil {
			panic(err)
		}
	}
}
