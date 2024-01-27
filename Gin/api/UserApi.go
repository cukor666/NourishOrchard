package api

import (
	"Gin/dao"
	"Gin/models"
	"Gin/response"
	"Gin/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserAdd 添加用户信息
func UserAdd(c *gin.Context) {
	var user models.User
	_ = c.ShouldBind(&user)
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

// UserList 查询所有用户信息，分页
func UserList(c *gin.Context) {
	var userDao dao.UserDao
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	log.Printf("pageSize = %v", pageSize)
	log.Printf("currentPage = %v", currentPage)
	users, result := userDao.SelectAllUserByLimit(pageSize, currentPage) // 新版，分页查询
	if result.Error != nil {
		log.Printf("select all by limit failed, err: %v\n", result.Error)
		response.Failed("后端查询数据失败", c)
	} else {
		response.Success("查询数据成功", gin.H{
			"users": users,
			"rows":  result.Rows,
		}, c)
	}
}

// FindUser 根据ID查找用户
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

// FindUserByStruct 通过表单的json转换为结构体对象，查询匹配的用户
func FindUserByStruct(c *gin.Context) {
	var (
		userDao dao.UserDao
		user    models.User
		users   []models.User
		rows    int64
		err     error
	)
	user.Promise = 1 // 只查普通用户
	user.Name = c.Query("name")
	user.NickName = c.Query("nickname")
	user.Gender = c.Query("gender")
	user.Address = c.Query("address")
	user.Phone = c.Query("phone")
	birthday := c.QueryArray("birthday[]") // 前端传过来的是数组，要这样拿
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	// log.Printf("前端传过来的数据")
	// user.Show()
	// log.Printf("birthday数组的长度：%v", len(birthday))

	if len(birthday) == 2 {
		log.Printf("len(birthday) == 2")
		users, rows, err = userDao.SelectByStructWithBirthday(user, currentPage, pageSize, birthday)
	} else {
		log.Printf("len(birthday) != 2")
		users, rows, err = userDao.SelectByStruct(user, currentPage, pageSize)
	}
	if err != nil {
		log.Printf("select by struct failed, err: %v\n", err)
		response.Failed("查询用户失败", c)
	} else {
		response.Success("查询用户成功", gin.H{
			"users": users,
			"rows":  rows,
		}, c)
	}
}

// UserLogin 用户登录，旧版的，自从使用了token之后就没有地方再调用这个函数了，新版登录看 AuthHandler
//func UserLogin(c *gin.Context) {
//	var user models.User
//	var ok bool
//	c.BindJSON(&user) // 拿JSON数据要这样拿
//	var userDao dao.UserDao
//	user, ok = userDao.SelectUserByNameAndPwd(user.Name, user.Password)
//	if !ok {
//		response.Failed("后端发送数据，用户名或密码错误", c)
//	} else {
//		response.Success("登录成功", user.Name, c)
//	}
//}

// AuthHandler 用户鉴权，用户登录时使用，新版的登录
func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user models.User
	err := c.ShouldBind(&user)
	// 校验用户名和密码是否正确
	if err != nil {
		log.Printf("should bind failed, err: %v\n", err)
		response.Failed("无效的参数", c)
		return
	}
	log.Println("user: ", user)
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
	tokenString, _ := utils.GenToken(u)
	response.Success("success", tokenString, c)
}

// SeeUserInfo 查看个人用户信息
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
	var user models.User
	c.BindJSON(&user) // 拿JSON数据要这样拿
	var userDao dao.UserDao
	ok := userDao.UpdateUserInfo(user)
	if !ok {
		response.Failed("更新失败，参数错误 ", c)
	} else {
		response.Success("更新用户信息成功", user, c)
	}
}

// 删除用户，根据ID
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("ID"))
	promise, _ := strconv.Atoi(c.Query("promise")) // 删除操作比较危险，只有权限大于普通用户的才可以删除
	if promise < 2 {
		response.Failed("删除失败，权限不够", c)
		return
	}
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

// 获取登录用户权限
func GetUserPromise(c *gin.Context) {
	name := c.Param("name")
	// log.Printf("name = %s", name)	// 能够获取到用户名
	var userDao dao.UserDao
	promise, ok := userDao.SelectPromiseByName(name)
	log.Printf("promise = %d", promise)
	if !ok {
		response.Failed("获取用户权限失败", c)
		return
	}
	response.Success("获取用户权限成功", promise, c)
}

// 获取用户个数
func GetUserCounnt(c *gin.Context) {
	s := c.Query("promise")
	promise, _ := strconv.Atoi(s)
	var userDao dao.UserDao
	userCount, err := userDao.SelectUserCount(int8(promise))
	if err != nil {
		log.Printf("get user count failed, err: %v\n", err)
		response.Failed("获取用户总条数失败", c)
		return
	}
	response.Success("获取用户总条数成功", userCount, c)
}
