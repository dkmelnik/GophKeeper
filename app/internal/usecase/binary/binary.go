package binary

import (
	"context"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

//go:generate mockgen -source=binary.go -destination=mocks/binary_repo_mock.go -package=mocks
type BinaryRepository interface {
	Save(ctx context.Context, user entity.Binary) (entity.Binary, error)
	FindByUserID(ctx context.Context, userID entity.ID) ([]entity.Binary, error)
}

// UseCase contains the business logic operations related to binary management.
type UseCase struct {
	binaryRepo BinaryRepository
}

// NewUseCase creates and initializes a new UseCase instance.
func NewUseCase(binaryRepo BinaryRepository) *UseCase {
	return &UseCase{binaryRepo}
}

func (uc *UseCase) Create(ctx context.Context, binary entity.Binary) (entity.Binary, error) {
	created, err := uc.binaryRepo.Save(ctx, binary)
	if err != nil {
		return entity.Binary{}, err
	}
	return created, err
}

func (uc *UseCase) List(ctx context.Context, userID entity.ID) ([]entity.Binary, error) {
	list, err := uc.binaryRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return list, nil
}
