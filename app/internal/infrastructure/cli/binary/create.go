package binary

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
	base64   string
	metadata string
)

func RegisterCreateBinaryCmd(command *cobra.Command) {
	createCmd.Flags().StringVarP(&base64, "base64", "b", "", "base64 content")
	createCmd.Flags().StringVarP(&metadata, "metadata", "m", "", "Some metadata")
	createCmd.MarkFlagRequired("base64")
	command.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create_binary",
	Short: "Create a new binary entry",
	Run: func(cmd *cobra.Command, args []string) {
		req := transfer.Request[entity.Binary]{
			JWT: configs.Client.Token,
			Data: entity.Binary{
				BinaryContent: base64,
				Metadata:      utils.StringToMap(metadata),
			},
		}

		var reply transfer.Reply[dto.BinaryDetails]

		err := rpc.ClientRPC.Call(domain.BinaryCreateMethod.String(), &req, &reply)
		if err != nil {
			log.Fatalf("%s: can't create binary entry", err.Error())
		}

		fmt.Printf("binary details:\n")
		fmt.Printf("  content: %s\n", reply.Data.BinaryContent)
		fmt.Printf("  metadata: %s\n", reply.Data.Metadata)
	},
}
