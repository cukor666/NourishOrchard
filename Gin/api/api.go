package api

import (
	"Gin/dao"
	"Gin/moudels"
	"Gin/response"
	"Gin/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加用户信息
func UserAdd(c *gin.Context) {
	var user moudels.User
	c.ShouldBind(&user)
	// user.Show()
	var userDao dao.UserDao
	_, affect := userDao.Insert(user)
	user.Password = "*" // 防止密码泄露
	if affect != 1 {
		response.Failed("添加用户失败", c)
	} else {
		response.Success("添加用户", user, c)
	}
}

// 查询所有用户信息
func UserList(c *gin.Context) {
	var userDao dao.UserDao
	result, err := userDao.SelectAllUser() // 新版
	if err != nil {
		log.Printf("selectAll failed, err: %v\n", err)
		response.Failed("后端查询数据失败", c)
	} else {
		response.Success("查询数据成功", result, c)
	}
}

// 根据ID查找用户
func FindUser(c *gin.Context) {
	val := c.Param("id")
	id, _ := strconv.Atoi(val)
	var userDao dao.UserDao
	u, err := userDao.SelectById(uint(id))
	if err != nil {
		log.Printf("SelectById failed, err: %v\n", err)
		response.Failed("查询用户失败", c)
	} else {
		response.Success("查询用户成功", u, c)
	}
}

// 通过表单的json转换为结构体对象，查询匹配的用户
func FindUserByStruct(c *gin.Context) {
	var (
		userDao dao.UserDao
		user    moudels.User
		users   []moudels.User
		err     error
	)
	user.Promise = 1 // 只查普通用户
	user.Name = c.Query("name")
	user.Gender = c.Query("gender")
	user.Address = c.Query("address")
	user.Phone = c.Query("phone")
	birthday := c.QueryArray("birthday[]") // 前端传过来的是数组，要这样拿
	log.Printf("前端传过来的数据")
	user.Show()
	log.Printf("birthday数组的长度：%v", len(birthday))

	if len(birthday) == 2 {
		log.Printf("len(birthday) == 2")
		users, err = userDao.SelectByStructWithBirthday(user, birthday)
	} else {
		log.Printf("len(birthday) != 2")
		users, err = userDao.SelectByStruct(user)
	}
	if err != nil {
		log.Printf("select by struct failed, err: %v\n", err)
		response.Failed("查询用户失败", c)
	} else {
		response.Success("查询用户成功", users, c)
	}
}

// 用户登录，旧版的，自从使用了token之后就没有地方再调用这个函数了，新版登录看 AuthHandler
func UserLogin(c *gin.Context) {
	var user moudels.User
	var ok bool
	c.BindJSON(&user) // 拿JSON数据要这样拿
	var userDao dao.UserDao
	user, ok = userDao.SelectUserByNameAndPwd(user.Name, user.Password)
	if !ok {
		response.Failed("后端发送数据，用户名或密码错误", c)
	} else {
		response.Success("登录成功", user.Name, c)
	}
}

// 用户鉴权，用户登录时使用，新版的登录
func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user moudels.User
	err := c.ShouldBind(&user)

	if err != nil {
		log.Printf("should bind failed, err: %v\n", err)
		response.Failed("无效的参数", c)
		return
	}

	// 校验用户名和密码是否正确
	// 通过查询数据库
	var userDao dao.UserDao
	if user.Name == "" || user.Password == "" {
		user.Name = c.GetString("name")
		user.Password = c.GetString("password")
	}
	u, ok := userDao.SelectUserByNameAndPwd(user.Name, user.Password)
	if !ok {
		response.Failed("鉴权失败", c)
		return
	}
	tokenString, _ := utils.GenToken(u.Name, u.Password)
	response.Success("success", tokenString, c)
}

// 查看个人用户信息
func SeeUserInfo(c *gin.Context) {
	name := c.Query("name")
	var userDao dao.UserDao
	user, ok := userDao.SelectByName(name)
	if !ok {
		response.Failed("无此用户: "+name, c)
	} else {
		response.Success("查询用户成功", user, c)
	}
}

// 更新用户信息
func UpdateUserInfo(c *gin.Context) {
	var user moudels.User
	c.BindJSON(&user) // 拿JSON数据要这样拿
	var userDao dao.UserDao
	ok := userDao.UpdateUserInfo(user)
	if !ok {
		response.Failed("更新失败，参数错误 ", c)
	} else {
		user.Password = "*"
		response.Success("更新用户信息成功", user, c)
	}
}

// 删除用户，根据ID
func DeleteUser(c *gin.Context) {
	s := c.Query("ID")
	id, _ := strconv.Atoi(s)
	var userDao dao.UserDao
	ok := userDao.DeleteById(uint(id))
	if !ok {
		response.Failed("删除失败，参数错误", c)
	} else {
		response.Success("删除成功", id, c)
	}
}

// 验证码
func Captcha(c *gin.Context) {
	utils.Captcha(c, 4) // 4个长度的验证码
}

// 验证，验证码
func Verify(c *gin.Context) {
	type request struct {
		Code string
	}
	var req request
	c.ShouldBind(&req)
	// log.Printf("code = %v", req.Code)
	if utils.CaptchaVerify(c, req.Code) {
		response.Success("验证码校验成功", 0, c)
	} else {
		response.Failed("验证码校验失败", c)
	}
}
