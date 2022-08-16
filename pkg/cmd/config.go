package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Config() *cli.Command {
	return &cli.Command{
		Name:        "config",
		Aliases:     []string{},
		Description: "Long config command description",
		Usage:       "Manage configuration for gd",
		Action: func(context *cli.Context) error {
			fmt.Println("Config 🛠")
			return nil
		},
	}
}
