package router

import "server/controller"

var (
	accountController  controller.AccountController
	loginController    controller.LoginController
	registerController controller.RegisterController
)

const (
	clientDomain = "http://localhost:5173"
)
