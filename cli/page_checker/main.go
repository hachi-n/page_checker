package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "page_checker",
		Usage: "page_checker [sub commands] [flags]",
		Description: "Page Check.",
		Commands: []*cli.Command{
			imgCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
