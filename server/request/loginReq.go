package request

type LoginRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required"`
	Promise  string `json:"promise" binding:"required"`
}
