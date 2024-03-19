package request

type ForgetPwdReq struct {
	Username string `json:"username" form:"username"`
	Promise  string `json:"promise" form:"promise"`
	Phone    string `json:"phone" form:"phone"`
	Email    string `json:"email" form:"email"`
	Code     string `json:"code" form:"code"`
	Password string `json:"password" form:"password"`
}
