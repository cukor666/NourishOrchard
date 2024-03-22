package request

// PromotionRequest 晋升管理员请求
type PromotionRequest struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Mark     string `json:"mark" form:"mark"`
}
