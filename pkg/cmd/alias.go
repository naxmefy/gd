package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Alias() *cli.Command {
	return &cli.Command{
		Name:        "alias",
		Aliases:     []string{},
		Description: "Long alias command description",
		Usage:       "Create custom command shortcuts",
		Action: func(context *cli.Context) error {
			fmt.Println("Alias 🛠")
			return nil
		},
	}
}
