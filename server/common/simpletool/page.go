package simpletool

// Page 用于分页查询
type Page struct {
	Size int `json:"pageSize" form:"pageSize" binding:"required,gt=0"`
	Num  int `json:"pageNum" form:"pageNum" binding:"required,gt=0"`
}
