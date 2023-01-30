package controllers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secretkey") // Need to replace with env variable

type Token struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email, role string) (string, error) {

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Token{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)

	return tokenString, err
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Token{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Token)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
