package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/errs"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/utils"
)

//go:generate mockgen -source=user.go -destination=mocks/user_repo_mock.go -package=mocks
type UserRepository interface {
	Save(ctx context.Context, user entity.User) (entity.User, error)
	IsEntryByLogin(ctx context.Context, login string) (bool, error)
	FindOneByLogin(ctx context.Context, login string) (entity.User, error)
}

//go:generate mockgen -source=user.go -destination=mocks/jwt_mock.go -package=mocks
type JWT interface {
	BuildJWTString(userID entity.ID, jti string) (string, error)
}

// UseCase contains the business logic operations related to user management.
type UseCase struct {
	bcryptCost int
	userRepo   UserRepository
	jwt        JWT
}

// NewUseCase creates and initializes a new UseCase instance.
func NewUseCase(userRepo UserRepository, jwt JWT) *UseCase {
	return &UseCase{bcrypt.DefaultCost, userRepo, jwt}
}

func (uc *UseCase) Register(ctx context.Context, user entity.User) (string, error) {
	exist, err := uc.userRepo.IsEntryByLogin(ctx, user.Login)
	if err != nil {
		return "", err
	}

	if exist {
		return "", fmt.Errorf("%w: login: %v", errs.ErrIsExist, user.Login)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), uc.bcryptCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	created, err := uc.userRepo.Save(ctx, user)
	if err != nil {
		return "", err
	}

	token, err := uc.jwt.BuildJWTString(created.ID, utils.GenerateGUID())

	return token, nil
}

func (uc *UseCase) Login(ctx context.Context, login, password string) (string, error) {
	user, err := uc.userRepo.FindOneByLogin(ctx, login)
	if err != nil {
		return "", err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errs.ErrInvalidCredentials
	}
	token, err := uc.jwt.BuildJWTString(user.ID, utils.GenerateGUID())
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc *UseCase) ChangePassword(ctx context.Context, prev, new string) error {
	return nil
}
