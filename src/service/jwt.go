package service

import (
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
	Role string
	jwt.RegisteredClaims
}

func (s *JwtService) GenerateToken(id, role string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			Subject:   id,
			Issuer:    "Erosync Ltd",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString(s.secret)
}

type JwtPayload struct {
	Sub  string
	Role string
}

func (s *JwtService) VerifyToken(tokenString string) (*JwtPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return s.secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("unable to covert token to claims")
	}

	return &JwtPayload{
		Sub:  claims.Subject,
		Role: claims.Role,
	}, nil
}
