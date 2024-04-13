package regctrl

import (
	"server/models"
	"server/utils/promisetool"
)

type registerResponse struct {
	Username string `json:"username"`
	Promise  string `json:"promise"`
}

func (res *registerResponse) set(account models.Account) {
	res.Username = account.Username
	res.Promise = promisetool.ToString(account.Promise)
}
