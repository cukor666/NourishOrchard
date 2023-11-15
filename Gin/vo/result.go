package vo

type Result struct {
	Rows  int64
	Error error
}

type DelteAdminType struct {
	ID      int `json: "ID"`
	Promise int `json: "promise"`
}
