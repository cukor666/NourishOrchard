package controller

import (
	"server/service"
)

type AccountController struct{}
type LoginController struct{}
type RegisterController struct{}
type UserController struct{}

var (
	accountService  service.AccountService
	loginService    service.LoginService
	registerService service.RegisterService
	userService     service.UserService
)
