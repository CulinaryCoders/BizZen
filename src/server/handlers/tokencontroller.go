package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"server/config"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// ? Official Go docs mentioned this package for managing JWTs (possibly more mature?): github.com/golang-jwt/jwt/v4

var mySigningKey = GetSigningKey()

// TODO: Add comment documentation (func GetSigningKey)
func GetSigningKey() []byte {
	// TODO: Config var or specific "JWT_KEY" value should probably be passed to this function instead of re-initializing entire config.
	config.InitEnvConfigs()
	var mySigningKey = []byte(config.ConfigVars.JWT_Key)
	return mySigningKey
}

// TODO: Add comment documentation (type Claims)
type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type TokenDetails struct {
	AccessToken         string
	RefreshToken        string
	AccessUuid          string
	RefreshUuid         string
	AccessTokenExpires  int64
	RefreshTokenExpires int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

// TODO: Add comment documentation (func GenerateToken)
func GenerateTokens(userid uint64, accountType string) (*TokenDetails, error) {

	// TODO: Add comment documentation (var expirationTime)
	expirationTime := time.Now().Add(time.Minute * 15).Unix()

	// Create Access Token
	tokenDetails := &TokenDetails{}
	tokenDetails.AccessTokenExpires = expirationTime
	tokenDetails.AccessUuid = uuid.NewV4().String()
	tokenDetails.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenDetails.RefreshUuid = uuid.NewV4().String()

	accessClaims := jwt.MapClaims{}
	accessClaims["authorized"] = true
	accessClaims["access_uuid"] = tokenDetails.AccessUuid
	accessClaims["user_id"] = userid
	accessClaims["exp"] = tokenDetails.AccessTokenExpires
	accessClaims["account_type"] = accountType

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)

	var err error
	tokenDetails.AccessToken, err = accessToken.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	refreshClaims := jwt.MapClaims{}
	refreshClaims["refresh_uuid"] = tokenDetails.RefreshUuid
	refreshClaims["user_id"] = userid
	refreshClaims["exp"] = tokenDetails.RefreshTokenExpires
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	tokenDetails.RefreshToken, err = refreshToken.SignedString(mySigningKey) // TO DO: Add different signing key
	if err != nil {
		return nil, err
	}

	return tokenDetails, err
}

func CreateToken(email, role string) (string, error) {

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

func ExtractToken(request *http.Request) string {
	bearToken := request.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// TODO: Add comment documentation (func ValidateToken)
func VerifyToken(request *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(request)
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	return token, err

}

func (h Handler) CreateAuth(userid uint64, tokenDetails *TokenDetails) error {
	accessToken := time.Unix(tokenDetails.AccessTokenExpires, 0)
	refreshToken := time.Unix(tokenDetails.RefreshTokenExpires, 0)
	now := time.Now()

	errAccess := h.client.Set(tokenDetails.AccessUuid, strconv.Itoa(int(userid)), accessToken.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := h.client.Set(tokenDetails.RefreshUuid, strconv.Itoa(int(userid)), refreshToken.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func TokenValid(request *http.Request) error {
	token, err := VerifyToken(request)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(request *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(request)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func (h Handler) FetchAuth(authD *AccessDetails) (uint64, error) {
	userid, err := h.client.Get(authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}
