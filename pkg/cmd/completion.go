package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Completion() *cli.Command {
	return &cli.Command{
		Name:        "completion",
		Aliases:     []string{},
		Description: "Long completion command description",
		Usage:       "Generate shell completion scripts",
		Action: func(context *cli.Context) error {
			fmt.Println("Completion 🛠")
			return nil
		},
	}
}
