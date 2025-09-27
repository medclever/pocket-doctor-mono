package messages

import (
	"context"
	"editor/app/types"
	"fmt"

	"github.com/urfave/cli/v3"
)

func listCommand(app types.App) *cli.Command {
	return &cli.Command{
		Name:        "ls",
		Description: "ls --lang=ru",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "lang",
				Value:    "ru",
				Usage:    "language for the greeting",
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			languageCode := cmd.String("lang")
			for _, message := range app.MessagesGetList().GetMessages() {
				fmt.Println(message.View(languageCode))
			}
			return nil
		},
	}
}
