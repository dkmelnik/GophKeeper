package text

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
	content  string
	metadata string
)

func RegisterCreateTextCmd(command *cobra.Command) {
	createCmd.Flags().StringVarP(&content, "content", "c", "", "Text content")
	createCmd.Flags().StringVarP(&metadata, "metadata", "m", "", "Some metadata")
	createCmd.MarkFlagRequired("content")
	command.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create_text",
	Short: "Create a new text entry",
	Run: func(cmd *cobra.Command, args []string) {
		req := transfer.Request[entity.Text]{
			JWT: configs.Client.Token,
			Data: entity.Text{
				TextContent: content,
				Metadata:    utils.StringToMap(metadata),
			},
		}

		var reply transfer.Reply[dto.TextDetails]

		err := rpc.ClientRPC.Call(domain.TextCreateMethod.String(), &req, &reply)
		if err != nil {
			log.Fatalf("%s: can't create text entry", err.Error())
		}

		fmt.Printf("text details:\n")
		fmt.Printf("  content: %s\n", reply.Data.TextContent)
		fmt.Printf("  metadata: %s\n", reply.Data.Metadata)
	},
}
