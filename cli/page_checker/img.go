package main

import (
	"github.com/hachi-n/page_checker/cli/page_checker/internal/options"
	checker "github.com/hachi-n/page_checker/lib/checker/img"
	"github.com/urfave/cli/v2"
)

func imgCommand() *cli.Command {
	var jsonPath string
	return &cli.Command{
		Name:        "img",
		Usage:       "page_checker img --json ${YOUR_JSON_PATH}",
		Description: "page image file connection confirmation.",
		Flags: []cli.Flag{
			options.JsonFlag(&jsonPath, true, ""),
		},
		Action: func(c *cli.Context) error {
			return checker.Apply(jsonPath)
		},
	}
}
