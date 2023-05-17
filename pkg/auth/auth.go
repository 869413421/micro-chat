package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtClaims struct {
	ID   uint64
	Name string
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(id uint64, name, key string) (string, error) {
	claims := JwtClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "micro-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}
