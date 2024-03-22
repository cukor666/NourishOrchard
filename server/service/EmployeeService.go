package service

import (
	"errors"
	"fmt"
	"regexp"
	"server/common/simpletool"
	"server/dao"
	"server/models"
	mc "server/models/code"
	"server/request"
	"server/utils"
)

// Info 根据账号获取员工信息
func (e EmployeeService) Info(username string) (emp models.Employee, err error) {
	r := regexp.MustCompile("CZKJ[0-9]{3,}")
	matchString := r.MatchString(username)
	if !matchString {
		return models.Employee{}, errors.New("账号不符合系统规范")
	}
	return dao.EmployeeDao{}.SelectByUsername(username)
}

// Update 更新员工信息
func (e EmployeeService) Update(employee models.Employee) error {
	_, err := dao.EmployeeDao{}.Update(employee)
	return err
}

// ListWithPage 查询员工信息，带有分页
func (e EmployeeService) ListWithPage(p simpletool.Page, employee models.Employee) ([]models.Employee, int64, error) {
	result, total, err := dao.EmployeeDao{}.ListWithPage(p, employee)
	if err != nil {
		levelLog("查询失败")
		return nil, 0, err
	}
	return result, total, nil
}

// ForgetPassword 忘记密码
func (e EmployeeService) ForgetPassword(req request.ForgetPwdReq) error {
	_, err := dao.EmployeeDao{}.SelectByUsernameAndPhone(req.Username, req.Phone)
	if err != nil {
		levelLog("查询员工失败")
		return err
	}
	// 对新密码加密
	pwd, err := utils.GetPwd(req.Password)
	if err != nil {
		levelLog("新密码加密失败")
		return err
	}
	err = dao.AccountDao{}.ChangePassword(req.Username, string(pwd))
	if err != nil {
		levelLog("修改密码错误")
		return err
	}
	return nil
}

// Promotion 员工晋升管理员
func (e EmployeeService) Promotion(req request.PromotionRequest) error {
	cnt := dao.AccountDao{}.GetCountByUsername(req.Username, mc.EMPLOYEE)
	if cnt == 0 {
		str := fmt.Sprintf("%s表中无数据username = %s", models.Account{}.TableName(), req.Username)
		levelLog(str)
		return errors.New(str)
	}
	err := dao.EmployeeDao{}.Promotion(req.Username, req.Name, req.Email, req.Mark)
	if err != nil {
		levelLog("晋升管理员失败")
		return err
	}
	return nil
}
