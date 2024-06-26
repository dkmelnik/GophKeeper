package binary

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dkmelnik/GophKeeper/configs"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/domain"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/rpc"
)

func RegisterListBinaryCmd(command *cobra.Command) {
	command.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list_binary",
	Short: "Get binary list",
	Run: func(cmd *cobra.Command, args []string) {
		req := transfer.Request[transfer.Empty]{
			JWT: configs.Client.Token,
		}

		var reply transfer.Reply[[]dto.BinaryDetails]

		err := rpc.ClientRPC.Call(domain.BinaryListMethod.String(), req, &reply)
		if err != nil {
			log.Fatalf("%s: can't get binary list", err.Error())
		}

		for _, item := range reply.Data {
			fmt.Printf("binary details:\n")
			fmt.Printf("  content: %s\n", item.BinaryContent)
			fmt.Printf("  metadata: %s\n", item.Metadata)
		}
	},
}
