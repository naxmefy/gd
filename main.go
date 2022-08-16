package main

import (
	"github.com/naxmefy/gd/pkg/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "gd",
		Usage:    "CLI helper for godot devs",
		Version:  "0.0.1",
		HideHelp: true,
		Commands: []*cli.Command{
			cmd.Alias(),
			cmd.Addon(),
			cmd.Completion(),
			cmd.Config(),
			cmd.Editor(),
			cmd.Extension(),
			cmd.Store(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
