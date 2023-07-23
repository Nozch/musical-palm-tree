package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

func main() {

	mySigningKey :=[]byte("bubbub")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "John Doe"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return
	}

	fmt.Println(tokenString)
}