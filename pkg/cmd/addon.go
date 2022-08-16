package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Addon() *cli.Command {
	return &cli.Command{
		Name:        "addon",
		Aliases:     []string{},
		Description: "Long addon command description",
		Usage:       "Manage addons for your project",
		Action: func(context *cli.Context) error {
			fmt.Println("Addon 🛠")
			return nil
		},
	}
}
