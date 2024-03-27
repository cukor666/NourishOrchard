package router

import "server/controller"

var (
	accountController  controller.AccountController
	loginController    controller.LoginController
	registerController controller.RegisterController
	userController     controller.UserController
	employeeController controller.EmployeeController
	adminController    controller.AdminController
	fruitController    controller.FruitController
)
