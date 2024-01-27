package router

import "server/controller"

var (
	accountController  controller.AccountController
	loginController    controller.LoginController
	registerController controller.RegisterController
	userController     controller.UserController
)
