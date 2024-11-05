package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type JWT struct {
	secret_key []byte
}

func New(secret string) *JWT {
	return &JWT{
		secret_key: []byte(secret),
	}
}

func (r *JWT) NewAccessToken(id string) (string, error) {
	claims := &TokenClaims{
		ID: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	str, err := token.SignedString(r.secret_key)
	if err != nil {
		return "", err
	}

	return str, nil
}

func (r *JWT) Parse(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return r.secret_key, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return "", fmt.Errorf("failed to type cast token")
	}

	return claims.ID, nil
}
