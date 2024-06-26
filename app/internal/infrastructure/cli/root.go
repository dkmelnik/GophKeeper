package cli

import (
	"github.com/spf13/cobra"

	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/binary"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/card"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/config"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/rpc"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/text"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/user"
)

var RootCmd = &cobra.Command{
	Use:   "client",
	Short: "A CLI for store personal information",
	Long:  `A Command Line Interface for interacting store logins, passwords, binary data and other personal information safely and securely.`,
}

func init() {
	config.RegisterChangeAddrCmd(RootCmd)

	rpc.InitClientRPC()

	user.RegisterCreateUserCmd(RootCmd)
	user.RegisterLoginUserCmd(RootCmd)

	text.RegisterCreateTextCmd(RootCmd)
	text.RegisterListTextCmd(RootCmd)

	card.RegisterCreateCardCmd(RootCmd)
	card.RegisterListCardCmd(RootCmd)

	binary.RegisterCreateBinaryCmd(RootCmd)
	binary.RegisterListBinaryCmd(RootCmd)
}

func Execute() error {
	return RootCmd.Execute()
}
