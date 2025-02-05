package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func IsJWTValid(tokenString, secret string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	return claims, ok
}
