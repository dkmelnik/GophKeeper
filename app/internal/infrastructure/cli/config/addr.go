package config

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/dkmelnik/GophKeeper/configs"
	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli/utils"
)

var addr string

func RegisterChangeAddrCmd(command *cobra.Command) {
	changeAddrCmd.Flags().StringVarP(&addr, "ADDR", "a", "", "server ADDR")
	command.AddCommand(changeAddrCmd)
}

var changeAddrCmd = &cobra.Command{
	Use:   "change_addr",
	Short: "Change configuration address",
	Long:  `Change configuration address for GophKeeper CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		configs.Client.ADDR = addr

		if err := utils.UpdateConfigFile(configs.Client); err != nil {
			log.Fatal(err)
		}
		fmt.Println("file updated successfully")
	},
}
