package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func EditorList() *cli.Command {
	return &cli.Command{
		Name:        "list",
		Aliases:     []string{},
		Description: "Long editor list command description",
		Usage:       "Print a list of all installed godot editors",
		Action: func(context *cli.Context) error {
			fmt.Println("EditorList 😮")
			return nil
		},
	}
}
