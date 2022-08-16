package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Extension() *cli.Command {
	return &cli.Command{
		Name:        "extension",
		Aliases:     []string{},
		Description: "Long extension command description",
		Usage:       "Manage gd extensions",
		Action: func(context *cli.Context) error {
			fmt.Println("Extension 🛠")
			return nil
		},
	}
}
