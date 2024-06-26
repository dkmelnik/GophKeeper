package rpc

import (
	"net/rpc"

	"github.com/dkmelnik/GophKeeper/internal/delivery"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/binary"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/card"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/text"
	"github.com/dkmelnik/GophKeeper/internal/delivery/rpc/user"
)

// Register registers RPC handlers for different endpoints using provided configurations and dependencies.
func Register(
	r *rpc.Server,
	userMW delivery.UserMiddleware,
	userUC delivery.UserUseCase,
	textUC delivery.TextUseCase,
	cardUC delivery.CardUseCase,
	binaryUC delivery.BinaryUseCase,
) error {
	// HANDLERS -----------------------
	userEnd := user.NewHandler(userUC)
	textEnd := text.NewHandler(userMW, textUC)
	cardEnd := card.NewHandler(userMW, cardUC)
	binaryEnd := binary.NewHandler(userMW, binaryUC)
	// HANDLERS -----------------------

	if err := r.Register(userEnd); err != nil {
		return err
	}

	if err := r.Register(textEnd); err != nil {
		return err
	}

	if err := r.Register(cardEnd); err != nil {
		return err
	}

	if err := r.Register(binaryEnd); err != nil {
		return err
	}

	return nil
}
