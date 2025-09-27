package ui

import (
	"context"
	"editor/app"
	"fmt"

	"github.com/rivo/tview"
	"github.com/urfave/cli/v3"
)

func New(app *app.App) *cli.Command {
	return &cli.Command{
		Name: "ui",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
			list := tview.NewList()
			for i := range 30 {
				list.AddItem(fmt.Sprint(i), "", '1', nil)
			}
			flex := tview.NewFlex().
				AddItem(list, 0, 1, true).
				AddItem(box, 0, 1, true)
			if err := tview.NewApplication().SetRoot(flex, true).Run(); err != nil {
				return err
			}
			return nil
		},
	}
}
