package controllers

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ? Official Go docs mentioned this package for managing JWTs (possibly more mature?): github.com/golang-jwt/jwt/v4

var mySigningKey = GetSigningKey()

// TODO: Add comment documentation (type Claims)
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// TODO: Add comment documentation (func GenerateToken)
func GenerateToken(email, role string) (string, error) {

	// TODO: Add comment documentation (var expirationTime)
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
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

// TODO: Add comment documentation (func ValidateToken)
func ValidateToken(signedToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return "", err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return "", err
	}

	if ok && token.Valid {
		return claims.Role, err
	}

	return "", err

}
