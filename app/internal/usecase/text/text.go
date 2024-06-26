package text

import (
	"context"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

//go:generate mockgen -source=text.go -destination=mocks/text_repo_mock.go -package=mocks
type TextRepository interface {
	Save(ctx context.Context, user entity.Text) (entity.Text, error)
	FindByUserID(ctx context.Context, userID entity.ID) ([]entity.Text, error)
}

// UseCase contains the business logic operations related to text management.
type UseCase struct {
	textRepo TextRepository
}

// NewUseCase creates and initializes a new UseCase instance.
func NewUseCase(textRepo TextRepository) *UseCase {
	return &UseCase{textRepo}
}

func (uc *UseCase) Create(ctx context.Context, text entity.Text) (entity.Text, error) {
	created, err := uc.textRepo.Save(ctx, text)
	if err != nil {
		return entity.Text{}, err
	}
	return created, err
}

func (uc *UseCase) List(ctx context.Context, userID entity.ID) ([]entity.Text, error) {
	list, err := uc.textRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return list, nil
}
