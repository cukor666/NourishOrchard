package request

type ForgetPwdReq struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Promise  string `json:"promise" form:"promise" binding:"required,promise"`
	Phone    string `json:"phone" form:"phone" binding:"omitempty,phone"`
	Email    string `json:"email" form:"email" binding:"omitempty,email"`
	Code     string `json:"code" form:"code" binding:"required"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=20"`
}
