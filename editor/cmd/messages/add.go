package messages

import (
	"context"
	"editor/app/types"
	"errors"

	"github.com/urfave/cli/v3"
)

func addCommand(app types.App) *cli.Command {
	return &cli.Command{
		Name:        "add",
		Description: "add --lang=ru <message text>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "lang",
				Value:    "ru",
				Usage:    "language for the greeting",
				Required: true,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() < 1 {
				return errors.New("Yor should receive message param")
			}
			languageCode := cmd.String("lang")
			err := app.MessagesAdd(languageCode, FullMessage(cmd))
			if err != nil {
				return err
			}
			return app.Persist()
		},
	}
}
