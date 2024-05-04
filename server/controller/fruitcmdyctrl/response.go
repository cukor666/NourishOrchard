package fruitcmdyctrl

import (
	"encoding/json"
	"server/common/levellog"
	"server/models"
)

type FruitCmdyResponse struct {
	ID    uint64   `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Imgs  []string `json:"imgs"`
	Desc  string   `json:"desc"`
}

func (fcr *FruitCmdyResponse) set(fc models.FruitCommodity) error {
	fcr.ID = fc.ID
	fcr.Name = fc.Name
	fcr.Price = fc.Price
	err := json.Unmarshal([]byte(*fc.Imgs), &fcr.Imgs)
	if err != nil {
		levellog.Controller("JSON Unmarshal error")
		return err
	}
	fcr.Desc = fc.Desc
	return nil
}
