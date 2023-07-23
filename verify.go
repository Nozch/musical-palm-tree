package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

func main() {
	mySigningKey := []byte("secret")

	tokenString := "your JWT token"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["user"], "is authorized")
	} else {
		fmt.Println(err)
	}
}
