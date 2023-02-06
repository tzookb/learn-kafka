package kafka

import (
	"fmt"

	"github.com/urfave/cli/v2" // imports as package "cli"
)

func GetCommand() cli.Command {
	return cli.Command{
		Name:    "kafka",
		Aliases: []string{"a"},
		Usage:   "kafka actions",
		Action: func(cCtx *cli.Context) error {
			fmt.Println("added task: ", cCtx.Args().First())
			return nil
		},
	}
}
