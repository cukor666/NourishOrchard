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

var (
	accountDao  dao.AccountDao
	userDao     dao.UserDao
	employeeDao dao.EmployeeDao
	adminDao    dao.AdminDao
)
