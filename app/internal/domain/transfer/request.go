package transfer

import (
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

// TokenProvider defines an interface for token management.
type TokenProvider interface {
	Token() string
	SetUserID(userID entity.ID)
}

// Empty represents an empty interface.
type Empty = interface{}

// Request represents a generic request structure.
type Request[T any] struct {
	JWT    string    // JWT holds the authentication token.
	UserID entity.ID // UserID represents the user ID associated with the request.
	Data   T         // Data holds the actual data payload of the request.
}

// Token returns the JWT token associated with the request.
func (r *Request[T]) Token() string {
	return r.JWT
}

// SetUserID sets the user ID associated with the request.
func (r *Request[T]) SetUserID(userID entity.ID) {
	r.UserID = userID
}
