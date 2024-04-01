package supctrl

import "server/models"

type supReq struct {
	Name          string `json:"name" binding:"required,min=3,max=50"`
	Address       string `json:"address" binding:"required,min=3,max=200"`
	ContactPerson string `json:"contactPerson" binding:"required,min=2,max=20"`
	Phone         string `json:"phone" binding:"required,phone"`
	Reputation    int    `json:"reputation" binding:"required,gt=0,lte=100"`
}

func (s *supReq) toSupplier() models.Supplier {
	return models.Supplier{
		ID:            0,
		Name:          s.Name,
		Address:       s.Address,
		ContactPerson: s.ContactPerson,
		Phone:         s.Phone,
		Reputation:    s.Reputation,
	}
}

type supUpdateReq struct {
	ID            int64  `json:"id" binding:"required,gt=0"`
	Name          string `json:"name" binding:"omitempty,min=3,max=50"`
	Address       string `json:"address" binding:"omitempty,min=3,max=200"`
	ContactPerson string `json:"contactPerson" binding:"omitempty,min=2,max=20"`
	Phone         string `json:"phone" binding:"omitempty,phone"`
	Reputation    int    `json:"reputation" binding:"omitempty,gt=0,lte=100"`
}

func (s *supUpdateReq) toSupplier() models.Supplier {
	return models.Supplier{
		ID:            s.ID,
		Name:          s.Name,
		Address:       s.Address,
		ContactPerson: s.ContactPerson,
		Phone:         s.Phone,
		Reputation:    s.Reputation,
	}
}
