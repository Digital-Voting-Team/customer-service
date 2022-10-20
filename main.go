package main

import (
	"os"

	"github.com/Digital-Voting-Team/customer-service/internal/cli"
)

func main() {
	//os.Setenv("KV_VIPER_FILE", "config.yaml")
	//os.Setenv("AUTH_SERVICE", "http://127.0.0.1:9110/jwt/login")
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
