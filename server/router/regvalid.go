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

	validations := map[string]func(fl validator.FieldLevel) bool{
		"username":    valid.UsernameValid,
		"promise":     valid.PromiseValid,
		"gender":      valid.GenderValid,
		"phone":       valid.PhoneValid,
		"birthday":    valid.BirthdayValid,
		"email":       valid.EmailValid,
		"empPosition": valid.EmpPosition,
		"address":     valid.AddressValid,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for key, validationFunc := range validations {
			if err := v.RegisterValidation(key, validationFunc); err != nil {
				panic(err)
			}
		}
	}
}
