package card

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dkmelnik/GophKeeper/configs"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/domain"
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/rpc"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/utils"
)

var (
	number   string
	expiry   string
	cvv      string
	metadata string
)

func RegisterCreateCardCmd(command *cobra.Command) {
	createCmd.Flags().StringVarP(&number, "number", "n", "", "Card number")
	createCmd.Flags().StringVarP(&expiry, "expiry", "e", "", "Card expiry")
	createCmd.Flags().StringVarP(&cvv, "cvv", "c", "", "Card cvv")
	createCmd.Flags().StringVarP(&metadata, "metadata", "m", "", "Some metadata")
	createCmd.MarkFlagRequired("number")
	createCmd.MarkFlagRequired("expiry")
	createCmd.MarkFlagRequired("cvv")
	command.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create_card",
	Short: "Create a new card entry",
	Run: func(cmd *cobra.Command, args []string) {
		req := transfer.Request[entity.Card]{
			JWT: configs.Client.Token,
			Data: entity.Card{
				CardNumber: number,
				ExpiryDate: expiry,
				CVV:        cvv,
				Metadata:   utils.StringToMap(metadata),
			},
		}

		var reply transfer.Reply[dto.CardDetails]

		err := rpc.ClientRPC.Call(domain.CardCreateMethod.String(), &req, &reply)
		if err != nil {
			log.Fatalf("%s: can't create card entry", err.Error())
		}

		fmt.Printf("card details:\n")
		fmt.Printf("  number: %s\n", reply.Data.CardNumber)
		fmt.Printf("  expiry: %s\n", reply.Data.CardNumber)
		fmt.Printf("  cvv: %s\n", reply.Data.CVV)
		fmt.Printf("  metadata: %s\n", reply.Data.Metadata)
	},
}
