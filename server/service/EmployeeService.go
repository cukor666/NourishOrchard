package service

import (
	"errors"
	"fmt"
	"regexp"
	"server/common/levellog"
	"server/common/simpletool"
	"server/models"
	mc "server/models/code"
	"server/request"
	"server/utils/pwdtool"
)

// Info 根据账号获取员工信息
func (e EmployeeService) Info(username string) (emp models.Employee, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Employee{}, errors.New("账号不符合系统规范")
	}
	return employeeDao.SelectByUsername(username)
}

// Update 更新员工信息
func (e EmployeeService) Update(employee models.Employee) error {
	_, err := employeeDao.Update(employee)
	return err
}

// ListWithPage 查询员工信息，带有分页
func (e EmployeeService) ListWithPage(p simpletool.Page, employee models.Employee) ([]models.Employee, int64, error) {
	result, total, err := employeeDao.ListWithPage(p, employee)
	if err != nil {
		levellog.Service("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// ForgetPassword 忘记密码
func (e EmployeeService) ForgetPassword(req request.ForgetPwdReq) error {
	_, err := employeeDao.SelectByUsernameAndPhone(req.Username, req.Phone)
	if err != nil {
		levellog.Service("查询员工失败")
		return err
	}
	// 对新密码加密
	pwd, err := pwdtool.GetPwd(req.Password)
	if err != nil {
		levellog.Service("新密码加密失败")
		return err
	}
	err = accountDao.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levellog.Service("修改密码错误")
		return err
	}
	return nil
}

// Promotion 员工晋升管理员
func (e EmployeeService) Promotion(req request.PromotionRequest) error {
	cnt := accountDao.GetCountByUsername(req.Username, mc.EMPLOYEE)
	if cnt == 0 {
		var am models.Account
		str := fmt.Sprintf("%s表中无数据username = %s", am.TableName(), req.Username)
		levellog.Service(str)
		return errors.New(str)
	}
	err := employeeDao.Promotion(req.Username, req.Name, req.Email, req.Mark)
	if err != nil {
		levellog.Service("晋升管理员失败")
		return err
	}
	return nil
}
