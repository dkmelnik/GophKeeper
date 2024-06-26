package card

import (
	"context"
	"errors"

	"github.com/dkmelnik/GophKeeper/internal/delivery"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/middleware"
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

// CardHandler handles RPC requests related to card operations.
type CardHandler struct {
	userMW delivery.UserMiddleware
	cardUC delivery.CardUseCase
}

// NewHandler creates and initializes a new CardHandler instance.
func NewHandler(userMW delivery.UserMiddleware, cardUC delivery.CardUseCase) *CardHandler {
	return &CardHandler{userMW, cardUC}
}

func (h *CardHandler) Create(payload *transfer.Request[entity.Card], reply *transfer.Reply[dto.CardDetails]) error {
	handler := middleware.Chain(
		h.createHandler,
		h.userMW.Auth,
	)

	return handler(payload, reply)
}

func (h *CardHandler) createHandler(args interface{}, reply interface{}) error {
	req, ok := args.(*transfer.Request[entity.Card])
	if !ok {
		return errors.New("invalid arguments")
	}

	req.Data.UserID = req.UserID

	created, err := h.cardUC.Create(context.Background(), req.Data)
	if err != nil {
		return err
	}

	*reply.(*transfer.Reply[dto.CardDetails]) = transfer.Reply[dto.CardDetails]{Data: dto.ToCardDetails(created)}

	return nil
}

func (h *CardHandler) List(payload *transfer.Request[any], reply *transfer.Reply[[]dto.CardDetails]) error {
	handler := middleware.Chain(
		h.listHandler,
		h.userMW.Auth,
	)

	return handler(payload, reply)
}

func (h *CardHandler) listHandler(args interface{}, reply interface{}) error {
	req, ok := args.(*transfer.Request[transfer.Empty])
	if !ok {
		return errors.New("invalid arguments")
	}

	list, err := h.cardUC.List(context.Background(), req.UserID)
	if err != nil {
		return err
	}

	out := make([]dto.CardDetails, 0, len(list))

	for _, item := range list {
		out = append(out, dto.ToCardDetails(item))
	}

	*reply.(*transfer.Reply[[]dto.CardDetails]) = transfer.Reply[[]dto.CardDetails]{Data: out}

	return nil
}
