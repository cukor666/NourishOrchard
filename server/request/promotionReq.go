package request

// PromotionRequest 晋升管理员请求
type PromotionRequest struct {
	Username string `json:"username" form:"username" binding:"required,username"`
	Password string `json:"password" form:"password" binding:"required,min=3,max=20"`
	Name     string `json:"name" form:"name" binding:"required,min=2,max=20"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Mark     string `json:"mark" form:"mark" binding:"max=100"`
}
