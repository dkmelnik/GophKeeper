package user

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dkmelnik/GophKeeper/configs"
	"github.com/dkmelnik/GophKeeper/internal/domain"
	"github.com/dkmelnik/GophKeeper/internal/domain/entity"
	"github.com/dkmelnik/GophKeeper/internal/domain/transfer"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/rpc"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/utils"
)

func RegisterLoginUserCmd(command *cobra.Command) {
	loginCmd.Flags().StringVarP(&login, "login", "l", "", "User login")
	loginCmd.Flags().StringVarP(&password, "password", "p", "", "User password")
	loginCmd.MarkFlagRequired("login")
	loginCmd.MarkFlagRequired("password")
	command.AddCommand(loginCmd)
}

var loginCmd = &cobra.Command{
	Use:   "login_user",
	Short: "Login user",
	Long:  `Login user with the provided login and password.`,
	Run: func(cmd *cobra.Command, args []string) {
		dto := &entity.User{
			Login:    login,
			Password: password,
		}

		var reply transfer.Reply[string]

		err := rpc.ClientRPC.Call(domain.UserLoginMethod.String(), dto, &reply)
		if err != nil {
			log.Fatalf("%s: can't login user", err.Error())
		}

		configs.Client.Token = reply.Data

		if err := utils.UpdateConfigFile(configs.Client); err != nil {
			log.Fatal(err)
		}

		fmt.Println("user login successfully")

	},
}
