package simpletool

// Page 用于分页查询
type Page struct {
	Size int `json:"pageSize" form:"pageSize"`
	Num  int `json:"pageNum" form:"pageNum"`
}
