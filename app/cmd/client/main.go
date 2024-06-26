package main

import (
	"log"

	"github.com/dkmelnik/GophKeeper/internal/infrastructure/cli"
)

var (
	version   string
	buildDate string
)

func main() {
	log.Printf("Version: %s\n", version)
	log.Printf("Build Date: %s\n", buildDate)
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
