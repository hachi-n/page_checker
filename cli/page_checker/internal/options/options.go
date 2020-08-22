package options

import "github.com/urfave/cli/v2"

func JsonFlag(destination *string, required bool, defaultValue string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:        "json",
		Value:       defaultValue,
		Usage:       "",
		Required:    required,
		Destination: destination,
	}
}

