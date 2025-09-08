package messages

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name: "ls",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			fmt.Println("messages ls command")
			return nil
		},
	}
}
