// Package jwt provides utilities for JWT (JSON Web Tokens) handling.
package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/errs"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

// Jwt represents a JSON Web Token generator and parser.
type Jwt struct {
	secret   string        // Secret key used for signing and verifying tokens
	tokenExp time.Duration // Token expiration duration
}

// NewJwt creates a new Jwt instance with the given secret and token expiration duration.
func NewJwt(secret string, exp time.Duration) *Jwt {
	return &Jwt{
		secret,
		exp,
	}
}

// BuildJWTString generates a new JWT string using the provided user ID and JTI (JWT ID).
// Returns the generated JWT string or an error if token generation fails.
func (j *Jwt) BuildJWTString(userID entity.ID, jti string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, transfer.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.tokenExp)),
		},
		SUB: userID,
		JTI: jti,
	})

	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parses the provided JWT string and validates its authenticity.
// Returns the parsed claims or an error if parsing or validation fails.
func (j *Jwt) ParseToken(tokenString string) (transfer.Claims, error) {
	claims := transfer.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("%w: unexpected signing method: %v", errs.ErrInvalidToken, t.Header["alg"])
			}
			return []byte(j.secret), nil
		})

	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, errs.ErrInvalidToken
	}

	return claims, nil
}
