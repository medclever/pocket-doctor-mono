package messages

import (
	"context"
	"editor/app"
	"errors"
	"strings"

	"github.com/urfave/cli/v3"
)

func addCommand(app *app.App) *cli.Command {
	return &cli.Command{
		Name:        "add",
		Description: "add <message text>",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() < 1 {
				return errors.New("Yor should receive message param")
			}
			message := []string{}
			for i := range cmd.Args().Len() {
				message = append(message, cmd.Args().Get(i))
			}
			app.MessagesAdd(strings.Join(message, " "))
			return app.Persist()
		},
	}
}
