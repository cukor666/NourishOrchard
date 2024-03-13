package statucode

// 全局状态码
const (
	NOTFOUNDACCOUNT = 4005 // 找不到对应的账号
	PASSWORDFAILED  = 4006 // 密码错误
	NOTFOUNDTOKEN   = 4009 // 无法获取token
	TOKENERROR      = 4010 // 前端传递的token和redis中的token不一致
	SYSTEMERR       = 5000 // 系统错误
)
