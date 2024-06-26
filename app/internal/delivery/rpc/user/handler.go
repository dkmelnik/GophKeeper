package user

import (
	"context"
	"time"

	"github.com/dkmelnik/GophKeeper/internal/delivery"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

// UserHandler handles RPC requests related to user operations.
type UserHandler struct {
	userUC delivery.UserUseCase
}

// NewHandler creates and initializes a new UserHandler instance.
func NewHandler(userUC delivery.UserUseCase) *UserHandler {
	return &UserHandler{userUC}
}

func (h *UserHandler) Register(payload dto.Register, reply *transfer.Reply[string]) error {
	ctx := context.Background()

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	if err := payload.Validate(); err != nil {
		return err
	}

	token, err := h.userUC.Register(ctxTimeout, payload.ToEntity())
	if err != nil {
		return err
	}

	*reply = transfer.Reply[string]{Data: token}

	return nil
}

func (h *UserHandler) Login(payload dto.Login, reply *transfer.Reply[string]) error {
	ctx := context.Background()

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err := payload.Validate(); err != nil {
		return err
	}

	token, err := h.userUC.Login(ctxTimeout, payload.Login, payload.Password)
	if err != nil {
		return err
	}

	*reply = transfer.Reply[string]{Data: token}

	return nil
}
