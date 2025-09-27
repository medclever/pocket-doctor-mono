package messages

import (
	"context"
	"editor/app/types"
	"errors"

	"github.com/urfave/cli/v3"
)

func editCommand(app types.App) *cli.Command {
	return &cli.Command{
		Name:        "edit",
		Description: "edit --lang=ru id=messageId <message text>",
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
			if cmd.Args().Len() < 1 {
				return errors.New("Yor should receive message param")
			}
			messageId := cmd.String("id")
			languageCode := cmd.String("lang")
			err := app.MessagesEdit(messageId, languageCode, FullMessage(cmd))
			if err != nil {
				return err
			}
			return app.Persist()
		},
	}
}
