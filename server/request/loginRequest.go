package request

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Promise  string `json:"promise"`
}