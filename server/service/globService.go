package service

import (
	"server/dao"
)

type AccountService struct{}
type LoginService struct{}
type RegisterService struct{}
type UserService struct{}
type EmployeeService struct{}
type AdminService struct{}
type FruitService struct{}

var (
	accountDao  dao.AccountDao
	userDao     dao.UserDao
	adminDao    dao.AdminDao
	employeeDao dao.EmployeeDao
	fruitDao    dao.FruitDao
)
