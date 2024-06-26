package rpc

import (
	"fmt"
	"net/rpc"

	"github.com/dkmelnik/GophKeeper/configs"
)

var ClientRPC *rpc.Client

func InitClientRPC() {
	client, err := rpc.Dial("tcp", configs.Client.ADDR)
	if err != nil {
		fmt.Printf("[ERROR] %s: check addr and connect to server, for change addr use change_addr command, %s \n", err.Error(), configs.Client.ADDR)
	}
	ClientRPC = client
}
