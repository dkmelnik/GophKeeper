package binary

import (
	"context"
	"errors"

	"github.com/dkmelnik/GophKeeper/internal/delivery"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/middleware"
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
)

// BinaryHandler handles RPC requests related to binary operations.
type BinaryHandler struct {
	userMW   delivery.UserMiddleware
	binaryUC delivery.BinaryUseCase
}

// NewHandler creates and initializes a new BinaryHandler instance.
func NewHandler(userMW delivery.UserMiddleware, binaryUC delivery.BinaryUseCase) *BinaryHandler {
	return &BinaryHandler{userMW, binaryUC}
}

func (h *BinaryHandler) Create(payload *transfer.Request[entity.Binary], reply *transfer.Reply[dto.BinaryDetails]) error {
	handler := middleware.Chain(
		h.createHandler,
		h.userMW.Auth,
	)

	return handler(payload, reply)
}

func (h *BinaryHandler) createHandler(args interface{}, reply interface{}) error {
	req, ok := args.(*transfer.Request[entity.Binary])
	if !ok {
		return errors.New("invalid arguments")
	}

	req.Data.UserID = req.UserID

	created, err := h.binaryUC.Create(context.Background(), req.Data)
	if err != nil {
		return err
	}

	*reply.(*transfer.Reply[dto.BinaryDetails]) = transfer.Reply[dto.BinaryDetails]{Data: dto.ToBinaryDetails(created)}

	return nil
}

func (h *BinaryHandler) List(payload *transfer.Request[any], reply *transfer.Reply[[]dto.BinaryDetails]) error {
	handler := middleware.Chain(
		h.listHandler,
		h.userMW.Auth,
	)

	return handler(payload, reply)
}

func (h *BinaryHandler) listHandler(args interface{}, reply interface{}) error {
	req, ok := args.(*transfer.Request[transfer.Empty])
	if !ok {
		return errors.New("invalid arguments")
	}

	list, err := h.binaryUC.List(context.Background(), req.UserID)
	if err != nil {
		return err
	}

	out := make([]dto.BinaryDetails, 0, len(list))

	for _, item := range list {
		out = append(out, dto.ToBinaryDetails(item))
	}

	*reply.(*transfer.Reply[[]dto.BinaryDetails]) = transfer.Reply[[]dto.BinaryDetails]{Data: out}

	return nil
}
