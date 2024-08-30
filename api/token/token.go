package token

import (
	"api/config"
	"errors"

	"github.com/golang-jwt/jwt"
)

type Claim struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

func ExtractClaimToken(stringToken string) (*Claim, error) {
	cfg := config.LoadConfig()
	token, err := jwt.ParseWithClaims(stringToken, &Claim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SIGNING_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ValidToken(tokenString string) (bool, error) {
	_, err := ExtractClaimToken(tokenString)
	if err != nil {
		return false, err
	}
	return true, nil
}
