package models

// Supplier 供应商
type Supplier struct {
	ID            int64  `json:"id" gorm:"primarykey" form:"id"`
	Name          string `json:"name" form:"name"`
	Address       string `json:"address" form:"address"`
	ContactPerson string `json:"contactPerson" form:"contactPerson"`
	Phone         string `json:"phone" form:"phone"`
	Reputation    int    `json:"reputation" form:"reputation"`
}

func (s *Supplier) TableName() string {
	return "suppliers"
}

func (s *Supplier) SetZero() (id int64, name, address, contactPerson, phone string, reputation int) {
	id, name, address, contactPerson, phone, reputation = s.ID, s.Name, s.Address, s.ContactPerson, s.Phone, s.Reputation
	s.ID, s.Name, s.Address, s.ContactPerson, s.Phone, s.Reputation = 0, "", "", "", "", 0
	return
}
