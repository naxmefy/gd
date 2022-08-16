package cmd

import (
	"github.com/urfave/cli/v2"
)

func Editor() *cli.Command {
	return &cli.Command{
		Name:        "editor",
		Aliases:     []string{},
		Description: "Long editor command description",
		Usage:       "Manage local installations of the godot editor",
		HideHelp:    true,
		Subcommands: []*cli.Command{
			EditorList(),
		},
	}
}
