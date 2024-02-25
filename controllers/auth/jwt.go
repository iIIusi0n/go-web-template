package auth

import (
	"api-server/data"
	"time"

	"api-server/config"
	"github.com/dgrijalva/jwt-go/v4"
)

type TokenClaims struct {
	UserID int    `json:"uid"`
	Name   string `json:"name"`
	Email  string `json:"email"`

	jwt.StandardClaims
}

func BuildNewToken(user data.User) (string, error) {
	claims := TokenClaims{
		UserID: user.ID,
		Name:   user.RealName,
		Email:  user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.ServerSecret))
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	claims := TokenClaims{}
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(config.ServerSecret), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &claims, key)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return &claims, nil
}
