package fruitdto

// Detail 水果详情响应
type Detail struct {
	FId            int    `json:"fId" gorm:"primarykey;column:fId" form:"fId"`
	FName          string `json:"fName" gorm:"column:fName" form:"fName"`
	FWater         int    `json:"fWater" gorm:"column:fWater" form:"fWater"`
	FSugar         int    `json:"fSugar" gorm:"column:fSugar" form:"fSugar"`
	FShelfLife     int    `json:"fShelfLife" gorm:"column:fShelfLife" form:"fShelfLife"`
	FOrigin        string `json:"fOrigin" gorm:"column:fOrigin" form:"fOrigin"`
	SId            int    `json:"sId" gorm:"column:sId" form:"sId"`
	SName          string `json:"sName" gorm:"column:sName" form:"sName"`
	SAddress       string `json:"sAddress" gorm:"column:sAddress" form:"sAddress"`
	SContactPerson string `json:"sContactPerson" gorm:"column:sContactPerson" form:"sContactPerson"`
	SPhone         string `json:"sPhone" gorm:"column:sPhone" form:"sPhone"`
	SReputation    int    `json:"sReputation" gorm:"column:sReputation" form:"sReputation"`
}
