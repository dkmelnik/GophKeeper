package user

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dkmelnik/GophKeeper/configs"
	"github.com/dkmelnik/GophKeeper/internal/delivery/dto"
	"github.com/dkmelnik/GophKeeper/internal/domain"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/rpc"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/utils"
)

var (
	login    string
	password string
	confirm  string
)

func RegisterCreateUserCmd(command *cobra.Command) {
	createCmd.Flags().StringVarP(&login, "login", "l", "", "User login")
	createCmd.Flags().StringVarP(&password, "password", "p", "", "User password")
	createCmd.Flags().StringVarP(&confirm, "confirm", "c", "", "Confirm password")
	createCmd.MarkFlagRequired("login")
	createCmd.MarkFlagRequired("password")
	createCmd.MarkFlagRequired("confirm")
	command.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create_user",
	Short: "Create a new user",
	Long:  `Create a new user with the provided login and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		dto := &dto.Register{
			Login:           login,
			Password:        password,
			ConfirmPassword: confirm,
		}

		var reply transfer.Reply[string]

		err := rpc.ClientRPC.Call(domain.UserRegisterMethod.String(), dto, &reply)
		if err != nil {
			log.Fatalf("%s: can't create user", err.Error())
		}

		configs.Client.Token = reply.Data

		if err := utils.UpdateConfigFile(configs.Client); err != nil {
			log.Fatal(err)
		}

		fmt.Println("user create successfully")
	},
}
