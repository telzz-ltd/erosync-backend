package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Role string
	jwt.RegisteredClaims
}

func GenerateToken(id, role string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Subject:   id,
			Issuer:    "Erosync Ltd",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString([]byte("my-insecure-jwt-key"))
}

func VerifyToken(tokenString string) (id, role string, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte("my-insecure-jwt-key"), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return id, role, errors.New("unable to covert token to claims")
	}

	return claims.Subject, claims.Role, nil
}
