package service

import (
	"server/config"
	"server/dao"
)

type AccountService struct{}
type LoginService struct{}
type RegisterService struct{}
type UserService struct{}
type EmployeeService struct{}
type AdminService struct{}

var redisDB = config.GetConfig().RedisConfig.GetClient()

var (
	accountDao  dao.AccountDao
	userDao     dao.UserDao
	employeeDao dao.EmployeeDao
	adminDao    dao.AdminDao
)
