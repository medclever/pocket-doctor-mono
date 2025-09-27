package messages

import (
	"context"
	"editor/app/types"
	"fmt"

	"github.com/urfave/cli/v3"
)

func listCommand(app types.App) *cli.Command {
	return &cli.Command{
		Name: "ls",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			for _, message := range app.MessagesGetList().GetMessages() {
				fmt.Println(message.View())
			}
			return nil
		},
	}
}
