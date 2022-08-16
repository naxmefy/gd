package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Store() *cli.Command {
	return &cli.Command{
		Name:        "store",
		Aliases:     []string{},
		Description: "Long config command description",
		Usage:       "Manage configuration for gd",
		Action: func(context *cli.Context) error {
			fmt.Println("CONFIG 🛠")
			return nil
		},
	}
}
