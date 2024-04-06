package response

type RegisterResponse struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
	Promise  string `json:"promise"`
}
