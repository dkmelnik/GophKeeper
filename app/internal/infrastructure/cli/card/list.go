package card

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

func RegisterListCardCmd(command *cobra.Command) {
	command.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list_card",
	Short: "Get card list",
	Run: func(cmd *cobra.Command, args []string) {
		req := transfer.Request[transfer.Empty]{
			JWT: configs.Client.Token,
		}

		var reply transfer.Reply[[]dto.CardDetails]

		err := rpc.ClientRPC.Call(domain.CardListMethod.String(), req, &reply)
		if err != nil {
			log.Fatalf("%s: can't get text list", err.Error())
		}

		for _, item := range reply.Data {
			fmt.Printf("card details:\n")
			fmt.Printf("  number: %s\n", item.CardNumber)
			fmt.Printf("  expiry: %s\n", item.CardNumber)
			fmt.Printf("  cvv: %s\n", item.CVV)
			fmt.Printf("  metadata: %s\n", item.Metadata)
		}
	},
}
