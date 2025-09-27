package messages

import (
	"context"
	"editor/app/types"

	"github.com/urfave/cli/v3"
)

func translateCommand(app types.App) *cli.Command {
	return &cli.Command{
		Name:        "translate",
		Description: "translate --lang=ru -id=<message_id> <message>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "lang",
				Value:    "ru",
				Usage:    "language for the greeting",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "id",
				Value:    "",
				Usage:    "message id",
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			languageCode := cmd.String("lang")
			messageId := cmd.String("id")
			message := FullMessage(cmd)
			err := app.MessagesTranslate(languageCode, messageId, message)
			if err != nil {
				return err
			}
			return app.Persist()
		},
	}
}
