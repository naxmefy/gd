package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "gd",
		Usage:    "CLI helper for godot devs",
		Version:  "0.0.1",
		Commands: []*cli.Command{},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
