package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/common/levellog"
	"server/common/simpletool"
	"server/config"
	cm "server/controller/args/claims"
	"server/controller/args/header"
	"server/models"
	mc "server/models/code"
	"server/response"
	"server/service"
	"server/utils"
	"strconv"
)

// List 查询水果列表
func (fc *FruitController) List(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	_, err = config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}

	var (
		p     simpletool.Page
		fruit models.Fruit
	)

	err1 := context.ShouldBindQuery(&p)
	err2 := context.ShouldBindQuery(&fruit)
	if err1 != nil || err2 != nil {
		levellog.Controller("绑定参数解构错误")
		response.Failed(context, "参数错误")
		return
	}

	fruits, total, err := service.FruitService{}.List(p, fruit)
	if err != nil {
		levellog.Controller(fmt.Sprintf("查询水果列表失败，fruits = %v", fruits))
		response.Failed(context, "查询水果列表失败")
		return
	}
	response.Success(context, gin.H{
		"fruits": fruits,
		"total":  total,
	}, "查询水果列表成功")
}

// Detail 水果详情
func (fc *FruitController) Detail(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	_, err = config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}

	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		levellog.Controller("转换失败")
		response.Failed(context, "转换失败")
		return
	}
	res, err := service.FruitService{}.Detail(id)
	if err != nil {
		levellog.Controller("水果详情请求接口错误")
		response.Failed(context, "水果详情请求接口错误")
		return
	}

	response.Success(context, res, "查询成功")
}

// Insert 添加水果接口
func (fc *FruitController) Insert(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	var fruit models.Fruit
	switch promise {
	case mc.USER:
		levellog.Controller(fmt.Sprintf("权限不够，当前权限为%d", promise))
		response.Failed(context, "权限不够")
	case mc.EMPLOYEE, mc.ADMIN:
		err = context.ShouldBindJSON(&fruit)
		if err != nil {
			levellog.Controller("绑定数据失败")
			response.Failed(context, "绑定数据失败")
			return
		}
		err = service.FruitService{}.Insert(fruit)
		if err != nil {
			levellog.Controller("添加水果失败")
			response.Failed(context, "添加水果失败")
			return
		}
		response.Success(context, 0, "添加成功")
	default:
		levellog.Controller(fmt.Sprintf("未知权限%d", promise))
		response.Failed(context, "未知权限")
	}
}

// Delete 删除水果信息
func (fc *FruitController) Delete(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	switch promise {
	case mc.USER:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够")
	case mc.EMPLOYEE, mc.ADMIN:
		id, err := strconv.Atoi(context.Query("id"))
		if err != nil {
			levellog.Controller("获取id失败")
			response.Failed(context, "参数错误")
			return
		}
		fruit, err := service.FruitService{}.Delete(id)
		if err != nil {
			levellog.Controller("删除失败")
			response.Failed(context, "系统错误删除失败")
			return
		}
		response.Success(context, fruit, "删除成功")
	default:
		levellog.Controller("未知权限")
		response.Failed(context, "未知权限")
	}
}

// Update 更新水果接口
func (fc *FruitController) Update(context *gin.Context) {
	// 解析token
	authorization := context.GetHeader(header.Authorization)
	token, err := GetToken(authorization) // 能走到这一步说明已经校验过了，所以这里不需要再进行校验
	if err != nil {
		levellog.Controller("获取token失败，请检查token是否过期")
		response.Failed(context, "获取token失败，请检查token是否过期")
		return
	}
	// 解析token，或token里面的内容
	claims, err := config.ParseAndVerifyJWT(token)
	if err != nil {
		levellog.Controller("解析token失败")
		response.Failed(context, "解析token失败")
		return
	}
	promise := utils.PromiseToInt(claims[cm.Promise].(string))
	var fruit models.Fruit
	switch promise {
	case mc.USER:
		levellog.Controller("权限不够")
		response.Failed(context, "权限不够")
	case mc.EMPLOYEE, mc.ADMIN:
		err = context.ShouldBindJSON(&fruit)
		if err != nil {
			levellog.Controller(fmt.Sprintf("数据绑定失败fruit = %v", fruit))
			response.Failed(context, "参数错误")
			return
		}
		err = service.FruitService{}.Update(fruit)
		if err != nil {
			levellog.Controller("更新失败")
			response.Failed(context, "更新失败")
			return
		}
		response.Success(context, 0, "更新成功")
	default:
		levellog.Controller("未知权限")
		response.Failed(context, "未知权限")
	}
}
