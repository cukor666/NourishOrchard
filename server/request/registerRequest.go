package request

import "time"

type RegisterRequest struct {
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Promise  int       `json:"promise"`
	Gender   string    `json:"gender"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Birthday time.Time `json:"birthday"`
}