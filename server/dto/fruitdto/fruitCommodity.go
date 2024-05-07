package fruitdto

import "server/models"

type FruitCommodity struct {
	models.FruitCommodity
	Origin string `json:"origin"`
}
