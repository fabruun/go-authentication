package application

import (
	"time"

	"github.com/fabruun/go-authentication/domain"
	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenGenerator struct{}

func (p *JwtTokenGenerator) GenerateToken(u *domain.User) {
	secretBytes := []byte(Secret)
	signingKey := jwt.SigningMethodHS256.Hash.New().Sum(secretBytes)

	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		Issuer:    Issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(signingKey)
	if err != nil {
		panic(err)
	}
	u.Token = signedString
}
