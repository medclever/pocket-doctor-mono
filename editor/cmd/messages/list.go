package messages

import (
	"context"
	"editor/app"
	"fmt"

	"github.com/urfave/cli/v3"
)

func listCommand(app *app.App) *cli.Command {
	return &cli.Command{
		Name: "ls",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			for _, message := range app.MessagesGetList() {
				fmt.Println(message.View())
			}
			return nil
		},
	}
}
