package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/sirupsen/logrus"
)

func ParseToken(accessToken, tokenKey string) (map[string]interface{}, bool, bool) {

	isValid := false
	isExpired := false

	token, err := jwt.ParseWithClaims(accessToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenKey), nil
	})

	if err != nil || token == nil {
		logrus.Error(err)
		isValid = false
	}

	sub := map[string]interface{}{}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	_ = json.Unmarshal([]byte(claims.Subject), &sub)

	isValid = ok && claims.Subject != ""
	isExpired = !token.Valid

	return sub, isValid, isExpired
}
