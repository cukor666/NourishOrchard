package code

// 全局状态码
const (
	NotFoundAccount = 4005 // 找不到对应的账号
	PasswordFailed  = 4006 // 密码错误
	TokenError      = 4010 // 前端传递的token和redis中的token不一致
	SystemError     = 5000 // 系统错误
)
