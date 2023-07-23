package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"testing"
	"time"
)

var mySigningKey = []byte("secret")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "nozch"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

func TestGenerateJWT(t *testing.T) {
	tokenString, err := GenerateJWT()
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user"] != "nozch" || claims["authorized"] != true {
			t.Fatalf("Token claims do not match expected values")
		}
	} else {
		t.Fatalf("Token is not valid: %v", err)
	}
}
