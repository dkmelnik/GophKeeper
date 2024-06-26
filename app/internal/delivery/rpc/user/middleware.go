package user

import (
	"errors"

	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/middleware"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

//go:generate mockgen -source=user.go -destination=mocks/jwt_mock.go -package=mocks
type JWT interface {
	ParseToken(tokenString string) (transfer.Claims, error)
}

type Middleware struct {
	jwt JWT
}

func NewMiddleware(jwt JWT) *Middleware {
	return &Middleware{jwt}
}

func (m *Middleware) Auth(next middleware.HandlerFunc) middleware.HandlerFunc {
	return func(args interface{}, reply interface{}) error {
		provider, ok := args.(transfer.TokenProvider)
		if !ok {
			return errors.New("invalid request type")
		}

		cl, err := m.jwt.ParseToken(provider.Token())
		if err != nil {
			return err
		}

		provider.SetUserID(cl.SUB)

		return next(provider, reply)
	}
}
