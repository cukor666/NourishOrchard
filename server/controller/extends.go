package controller

import (
	"github.com/golang-jwt/jwt/v5"
	cm "server/controller/args/claims"
	"server/utils/promisetool"
)

func GetPromise(claims jwt.MapClaims) int {
	return promisetool.ToInt(claims[cm.Promise].(string))
}

func GetUsername(claims jwt.MapClaims) string {
	return claims[cm.Username].(string)
}
