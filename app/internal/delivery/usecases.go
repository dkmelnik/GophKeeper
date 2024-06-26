package delivery

import (
	"context"

	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/middleware"
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

//go:generate mockgen -source=text.go -destination=mocks/user_mw_mock.go -package=mocks
type UserMiddleware interface {
	Auth(next middleware.HandlerFunc) middleware.HandlerFunc
}

//go:generate mockgen -source=user.go -destination=mocks/user_uc_mock.go -package=mocks
type UserUseCase interface {
	Register(ctx context.Context, user entity.User) (string, error)
	Login(ctx context.Context, login, password string) (string, error)
	ChangePassword(ctx context.Context, prev, new string) error
}

//go:generate mockgen -source=text.go -destination=mocks/text_uc_mock.go -package=mocks
type TextUseCase interface {
	Create(ctx context.Context, text entity.Text) (entity.Text, error)
	List(ctx context.Context, userID entity.ID) ([]entity.Text, error)
}

//go:generate mockgen -source=text.go -destination=mocks/text_uc_mock.go -package=mocks
type CardUseCase interface {
	Create(ctx context.Context, text entity.Card) (entity.Card, error)
	List(ctx context.Context, userID entity.ID) ([]entity.Card, error)
}

//go:generate mockgen -source=handler.go -destination=mocks/binary_uc_mock.go -package=mocks
type BinaryUseCase interface {
	Create(ctx context.Context, binary entity.Binary) (entity.Binary, error)
	List(ctx context.Context, userID entity.ID) ([]entity.Binary, error)
}
