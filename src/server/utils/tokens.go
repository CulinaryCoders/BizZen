package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TODO: Add comment documentation (type Claims)
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

// TODO: Add comment documentation (func GenerateToken)
func GenerateToken(email string, role string, signingKey []byte) (string, error) {

	// Set token expiration time to one hour from current time
	expirationTime := jwt.NewNumericDate(time.Now().Add(1 * time.Hour))

	claims := Claims{
		email,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims.RegisteredClaims)
	tokenString, err := token.SignedString(signingKey)

	return tokenString, err
}

// TODO: Add comment documentation (func ValidateToken)
func ValidateToken(signedToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
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

	if claims.VerifyExpiresAt(time.Now(), true) {
		err = errors.New("token expired")
		return "", err
	}

	if ok && token.Valid {
		return claims.Role, err
	}

	return "", err

}
