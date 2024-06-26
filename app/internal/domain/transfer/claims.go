package transfer

import (
	"github.com/golang-jwt/jwt/v4"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

// Claims represents custom JWT claims.
type Claims struct {
	jwt.RegisteredClaims           // Embedded struct to inherit standard JWT claims
	SUB                  entity.ID // Subject (SUB) identifies the principal that is the subject of the JWT
	JTI                  string    // JWT ID (JTI) uniquely identifies the JWT
}
