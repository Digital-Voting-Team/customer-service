package main

import (
	"os"

	"customer-service/internal/cli"
)

func main() {
	os.Setenv("KV_VIPER_FILE", "config.yaml")
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
