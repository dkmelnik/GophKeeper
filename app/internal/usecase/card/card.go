package card

import (
	"context"

	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
)

//go:generate mockgen -source=card.go -destination=mocks/card_repo_mock.go -package=mocks
type CardRepository interface {
	Save(ctx context.Context, user entity.Card) (entity.Card, error)
	FindByUserID(ctx context.Context, userID entity.ID) ([]entity.Card, error)
}

// UseCase contains the business logic operations related to card management.
type UseCase struct {
	cardRepo CardRepository
}

// NewUseCase creates and initializes a new UseCase instance.
func NewUseCase(cardRepo CardRepository) *UseCase {
	return &UseCase{cardRepo}
}

func (uc *UseCase) Create(ctx context.Context, card entity.Card) (entity.Card, error) {
	created, err := uc.cardRepo.Save(ctx, card)
	if err != nil {
		return entity.Card{}, err
	}
	return created, err
}

func (uc *UseCase) List(ctx context.Context, userID entity.ID) ([]entity.Card, error) {
	list, err := uc.cardRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return list, nil
}
