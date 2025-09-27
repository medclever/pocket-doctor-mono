package ui

import (
	"context"
	"editor/app/types"
	"editor/app/utils"

	"github.com/rivo/tview"
	"github.com/urfave/cli/v3"
)

func New(app types.App) *cli.Command {
	return &cli.Command{
		Name:        "ui",
		Description: "ui --lang=ru",
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
			list := tview.NewList()
			for i, message := range app.MessagesGetList().GetMessages() {
				position := utils.IntToRune(i + 1)
				title := message.View(languageCode)
				meta := message.Meta(languageCode)
				list.AddItem(title, meta, position, nil)
			}
			if err := tview.NewApplication().SetRoot(list, true).Run(); err != nil {
				return err
			}
			return nil
		},
	}
}
