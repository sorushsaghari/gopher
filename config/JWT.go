package config

import "github.com/dgrijalva/jwt-go"

var JwtKey = []byte("my_secret_key")


type Credentials struct {
	Password string `json:"password"`
	Email string `json:"email"`
}

type JWT struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func ParseToken(tknstr string) (*JWT, error) {
	claims := &JWT{}

	token, err := jwt.ParseWithClaims(tknstr, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if !token.Valid {
		return nil, err
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}
		return nil, err
	}
	return claims, nil
}