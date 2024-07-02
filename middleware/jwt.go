package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("megh")

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(username, email string) (string, error) {
	claims := &Claims{
		Username:       username,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}
