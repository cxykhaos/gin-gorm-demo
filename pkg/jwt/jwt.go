package jwt

import (
	"FileCloud/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey []byte

type Claims struct {
	User_Id string
	jwt.StandardClaims
}

func ReleaseToken(user *model.User) (string, error) {
	expirationTime := time.Now().Add(24 * 7 * time.Hour)
	claims := &Claims{
		User_Id: user.User_Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "cxykhaos",
			Subject:   "user token",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	return token, claims, err
}
