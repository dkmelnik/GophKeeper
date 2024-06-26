package text

import (
	"context"
	"errors"

	"github.com/dkmelnik/GophKeeper/internal/delivery"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/middleware"
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

// TextHandler handles RPC requests related to text operations.
type TextHandler struct {
	userMW delivery.UserMiddleware
	textUC delivery.TextUseCase
}

// NewHandler creates and initializes a new TextHandler instance.
func NewHandler(userMW delivery.UserMiddleware, textUC delivery.TextUseCase) *TextHandler {
	return &TextHandler{userMW, textUC}
}

func (h *TextHandler) Create(payload *transfer.Request[entity.Text], reply *transfer.Reply[dto.TextDetails]) error {
	handler := middleware.Chain(
		h.createHandler,
		h.userMW.Auth,
	)

	return handler(payload, reply)
}

func (h *TextHandler) createHandler(args interface{}, reply interface{}) error {
	req, ok := args.(*transfer.Request[entity.Text])
	if !ok {
		return errors.New("invalid arguments")
	}

	req.Data.UserID = req.UserID

	created, err := h.textUC.Create(context.Background(), req.Data)
	if err != nil {
		return err
	}

	*reply.(*transfer.Reply[dto.TextDetails]) = transfer.Reply[dto.TextDetails]{Data: dto.ToTextDetails(created)}

	return nil
}

func (h *TextHandler) List(payload *transfer.Request[any], reply *transfer.Reply[[]dto.TextDetails]) error {
	handler := middleware.Chain(
		h.listHandler,
		h.userMW.Auth,
	)

	return handler(payload, reply)
}

func (h *TextHandler) listHandler(args interface{}, reply interface{}) error {
	req, ok := args.(*transfer.Request[transfer.Empty])
	if !ok {
		return errors.New("invalid arguments")
	}

	list, err := h.textUC.List(context.Background(), req.UserID)
	if err != nil {
		return err
	}
	out := make([]dto.TextDetails, 0, len(list))

	for _, item := range list {
		out = append(out, dto.ToTextDetails(item))
	}
	*reply.(*transfer.Reply[[]dto.TextDetails]) = transfer.Reply[[]dto.TextDetails]{Data: out}

	return nil
}
