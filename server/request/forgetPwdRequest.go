package request

type ForgetPwdReq struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Promise  string `json:"promise" form:"promise" binding:"required,promise"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Code     string `json:"code" form:"code" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
