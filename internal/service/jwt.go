package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secret []byte
}

func NewJwtService(secret string) *JwtService {
	return &JwtService{
		secret: []byte(secret),
	}
}

type Claims struct {
	jwt.RegisteredClaims
}

func (s *JwtService) GenerateToken(ctx context.Context, sub string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,
			Issuer:    "Erosync",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	})

	return token.SignedString(s.secret)
}

func (s *JwtService) ValidateToken(ctx context.Context, tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims.Subject, nil
	}

	return "", errors.New("unknown claims type, cannot proceed")
}
